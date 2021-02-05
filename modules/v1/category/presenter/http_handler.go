package presenter

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"pretest-privyid/helper"
	"pretest-privyid/modules/v1/category/model"
	"pretest-privyid/modules/v1/category/usecase"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */
type HTTPCategoryHandler struct {
	CategoryUsecase usecase.CategoryUsecase
}

func NewHTTPHandler(CategoryUsecase usecase.CategoryUsecase) *HTTPCategoryHandler {
	return &HTTPCategoryHandler{
		CategoryUsecase: CategoryUsecase,
	}
}

func (h *HTTPCategoryHandler) MountCategory(group *echo.Group)  {
	group.POST("/category", h.CreateCategory)
	group.GET("/categories", h.GetAllCategories)
	group.GET("/category/:id", h.GetCategoryById)
	group.PUT("/category/:id", h.UpdateCategory)
	group.DELETE("/category/:id", h.DeleteCategory)
}

func (h *HTTPCategoryHandler) CreateCategory(c echo.Context) error  {
	reqData := model.Category{}
	err := c.Bind(&reqData)
	if err != nil {
		log.Println("bind_request_data", err.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Parameter body kosong", nil))
	}

	saveResult := h.CategoryUsecase.CreateCategory(reqData)
	if saveResult.Error != nil {
		err := fmt.Errorf("Gagal menambah data category")
		log.Println(saveResult.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	data, ok := saveResult.Result.(model.Category)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		return c.JSON(http.StatusOK, helper.ResponseDetailOutput(err.Error(), data))
	}
	return c.JSON(http.StatusCreated, data)
}

func (h *HTTPCategoryHandler) GetAllCategories(c echo.Context) error {
	categories := h.CategoryUsecase.GetAllCategories()
	if categories.Error != nil {
		err := fmt.Errorf("Gagal mendapatkan category")
		log.Println(categories.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	result, _ := categories.Result.(model.Categories)
	return c.JSON(http.StatusOK, result)
}

func (h *HTTPCategoryHandler) GetCategoryById(c echo.Context) error {
	idCategory := c.Param("id")
	category := h.CategoryUsecase.GetCategoryById(idCategory)
	if category.Error != nil {
		err := fmt.Errorf("Gagal mendapatkan category")
		log.Println(category.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	result, _ := category.Result.(model.Categories)
	return c.JSON(http.StatusOK, result)
}

func (h *HTTPCategoryHandler) UpdateCategory(c echo.Context) error  {
	idCategory := c.Param("id")
	if idCategory == "" {
		log.Println("no_parameter")
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Parameter id diperlukan", nil))
	}
	reqData := model.Category{}
	err := c.Bind(&reqData)
	if err != nil {
		log.Println("bind_request_data", err.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Parameter body kosong", nil))
	}

	saveResult := h.CategoryUsecase.UpdateCategoryById(idCategory, reqData)
	if saveResult.Error != nil {
		err := fmt.Errorf("Gagal menambah data category")
		log.Println(saveResult.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	data, ok := saveResult.Result.(model.Category)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		return c.JSON(http.StatusOK, helper.ResponseDetailOutput(err.Error(), data))
	}
	return c.JSON(http.StatusCreated, data)
}

func (h *HTTPCategoryHandler) DeleteCategory(c echo.Context) error {
	idCategory := c.Param("id")
	if idCategory == "" {
		log.Println("no_parameter")
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput("Parameter id diperlukan", nil))
	}
	deleteResult := h.CategoryUsecase.DeleteCategory(idCategory)
	if deleteResult.Error != nil {
		err := fmt.Errorf("Gagal menghapus data category")
		log.Println(deleteResult.Error.Error())
		return c.JSON(http.StatusBadRequest, helper.ResponseDetailOutput(err.Error(), nil))
	}
	result := echo.Map{"message" : "Data Category berhasil dihapus"}
	return c.JSON(http.StatusOK, result)
}