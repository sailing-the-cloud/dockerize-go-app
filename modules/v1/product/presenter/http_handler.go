package presenter

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"log"
	"net/http"
	"os"
	"pretest-privyid/helper"
	"pretest-privyid/modules/v1/product/model"
	"pretest-privyid/modules/v1/product/usecase"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */
type HTTPProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewHTTPHandler(ProductUsecase usecase.ProductUsecase) *HTTPProductHandler {
	return &HTTPProductHandler{
		ProductUsecase: ProductUsecase,
	}
}

func (h *HTTPProductHandler) MountProduct(group *echo.Group)  {
	group.POST("/product", h.CreateProduct)
	group.POST("/product-image", h.UploadImage)
	group.GET("/products", h.GetAllProduct)
	group.GET("/product/:id", h.GetProductById)
	group.PUT("/product/:id", h.UpdateProduct)
	group.DELETE("/product/:id", h.DeleteProduct)
}

func (h *HTTPProductHandler) CreateProduct(c echo.Context) error  {
	reqData := model.Product{}
	err := c.Bind(&reqData)
	if err != nil {
		log.Println("bind_request_data", err.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Parameter body kosong", nil))
	}

	saveResult := h.ProductUsecase.CreateProduct(reqData)
	if saveResult.Error != nil {
		err := fmt.Errorf("Gagal menambah data category")
		log.Println(saveResult.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	data, ok := saveResult.Result.(model.Product)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		return c.JSON(http.StatusOK, helper.ResponseDetailOutput(err.Error(), data))
	}
	return c.JSON(http.StatusCreated, data)
}

func (h *HTTPProductHandler) UploadImage(c echo.Context) error {
	productId := c.FormValue("productId")
	imageName := c.FormValue("imageName")
	productImage, errorFile := c.FormFile("productImage")
	if errorFile != nil {
		err := fmt.Errorf("File kosong")
		log.Println(errorFile.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	src, error := productImage.Open()
	if error != nil {
		err := fmt.Errorf("Gagal membaca file")
		log.Println(error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	defer src.Close()

	// Destination
	path := "images/" + productImage.Filename
	dst, error := os.Create(path)
	if error != nil {
		err := fmt.Errorf("Gagal menyimpan file")
		log.Println(error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	defer dst.Close()

	// Copy
	if _, error = io.Copy(dst, src); error != nil {
		err := fmt.Errorf("Gagal menyimpan file")
		log.Println(error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}

	image := model.Image{}
	image.Name = imageName
	image.File = path

	saveResult := h.ProductUsecase.UploadImage(productId, image)
	if saveResult.Error != nil {
		err := fmt.Errorf("Gagal upload foto product")
		log.Println(saveResult.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	result := echo.Map{"message" : "Berhasil upload gambar"}
	return c.JSON(http.StatusCreated, result)
}

func (h *HTTPProductHandler) GetAllProduct(c echo.Context) error {
	products := h.ProductUsecase.GetAllProduct()
	if products.Error != nil {
		err := fmt.Errorf("Gagal mendapatkan product")
		log.Println(products.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	result, _ := products.Result.([]model.Product)
	return c.JSON(http.StatusOK, result)
}

func (h *HTTPProductHandler) GetProductById(c echo.Context) error {
	productId := c.Param("id")
	products := h.ProductUsecase.GetProductById(productId)
	if products.Error != nil {
		err := fmt.Errorf("Gagal mendapatkan product")
		log.Println(products.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	result, _ := products.Result.([]model.Product)
	if result == nil {
		err := fmt.Errorf("Product sudah dihapus")
		log.Println(products.Error.Error())
		return c.JSON(http.StatusNoContent, helper.ResponseDetailOutput(err.Error(), nil))
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTPProductHandler) UpdateProduct(c echo.Context) error {
	productId := c.Param("id")

	reqData := model.Product{}
	err := c.Bind(&reqData)
	if err != nil {
		log.Println("bind_request_data", err.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Parameter body kosong", nil))
	}

	saveResult := h.ProductUsecase.UpdateProduct(productId, reqData)
	if saveResult.Error != nil {
		err := fmt.Errorf("Gagal menambah data category")
		log.Println(saveResult.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	data, ok := saveResult.Result.(model.Product)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		return c.JSON(http.StatusOK, helper.ResponseDetailOutput(err.Error(), data))
	}
	return c.JSON(http.StatusCreated, data)
}

func (h *HTTPProductHandler) DeleteProduct(c echo.Context) error {
	idProduct := c.Param("id")
	if idProduct == "" {
		log.Println("no_parameter")
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Parameter id diperlukan", nil))
	}
	deleteResult := h.ProductUsecase.DeleteProduct(idProduct)
	if deleteResult.Error != nil {
		err := fmt.Errorf("Gagal menghapus data product")
		log.Println(deleteResult.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	result := echo.Map{"message" : "Data Product berhasil dihapus"}
	return c.JSON(http.StatusOK, result)
}