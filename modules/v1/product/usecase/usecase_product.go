package usecase

import (
	"fmt"
	"log"
	model2 "pretest-privyid/modules/v1/category/model"
	"pretest-privyid/modules/v1/product/model"
	"pretest-privyid/modules/v1/product/repository"
	"strconv"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type ProductUsecaseImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductUseCase(ProductRepo repository.ProductRepository) ProductUsecase {
	return &ProductUsecaseImpl{
		ProductRepository: ProductRepo,
	}
}

func (rp *ProductUsecaseImpl) CreateProduct(param model.Product) ResultUseCase {
	output := ResultUseCase{}

	saveResult := rp.ProductRepository.CreateProduct(param)
	if saveResult.Error != nil {
		log.Println("Error menyimpan data :", saveResult.Error.Error())
		output = ResultUseCase{Error: saveResult.Error}
		return output
	}
	data, ok := saveResult.Result.(model.Product)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		output = ResultUseCase{Error: err}
		return output
	}
	output = ResultUseCase{Result: data}
	return output
}

func (rp *ProductUsecaseImpl) UploadImage(productId string, param model.Image) ResultUseCase {
	output := ResultUseCase{}

	uploadResult := rp.ProductRepository.UploadImage(productId, param)
	if uploadResult.Error != nil {
		log.Println("Error menyimpan gambar :", uploadResult.Error.Error())
		output = ResultUseCase{Error: uploadResult.Error}
		return output
	}
	output = ResultUseCase{Result: param}
	return output
}

func (rp *ProductUsecaseImpl) GetAllProduct() ResultUseCase {
	output := ResultUseCase{}

	var products []model.Product

	resultProduct := rp.ProductRepository.GetAllProduct()
	if resultProduct.Error != nil {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		output = ResultUseCase{Error: err}
		return output
	}
	products = resultProduct.Result.([]model.Product)

	for i := 0; i < len(products); i++ {
		id := strconv.Itoa(products[i].ID)
		resultCategory := rp.ProductRepository.GetCategoryOfProduct(id)
		if resultCategory.Error != nil {
			err := fmt.Errorf("Gagal mendapatkan data category")
			log.Println(err.Error())
			output = ResultUseCase{Error: err}
			return output
		}
		category, _ := resultCategory.Result.([]model2.Category)
		products[i].Categories = category
	}

	for i := 0; i < len(products); i++ {
		id := strconv.Itoa(products[i].ID)
		resultCategory := rp.ProductRepository.GetImageOfProduct(id)
		if resultCategory.Error != nil {
			err := fmt.Errorf("Gagal mendapatkan data category")
			log.Println(err.Error())
			output = ResultUseCase{Error: err}
			return output
		}
		images, _ := resultCategory.Result.([]model.Image)
		products[i].Image = images
	}

	output = ResultUseCase{Result: products}
	return output
}

func (rp *ProductUsecaseImpl) GetProductById(productId string) ResultUseCase {
	output := ResultUseCase{}

	var products []model.Product

	resultProduct := rp.ProductRepository.GetProductById(productId)
	if resultProduct.Error != nil {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		output = ResultUseCase{Error: err}
		return output
	}
	products = resultProduct.Result.([]model.Product)

	for i := 0; i < len(products); i++ {
		id := strconv.Itoa(products[i].ID)
		resultCategory := rp.ProductRepository.GetCategoryOfProduct(id)
		if resultCategory.Error != nil {
			err := fmt.Errorf("Gagal mendapatkan data category")
			log.Println(err.Error())
			output = ResultUseCase{Error: err}
			return output
		}
		category, _ := resultCategory.Result.([]model2.Category)
		products[i].Categories = category
	}

	for i := 0; i < len(products); i++ {
		id := strconv.Itoa(products[i].ID)
		resultCategory := rp.ProductRepository.GetImageOfProduct(id)
		if resultCategory.Error != nil {
			err := fmt.Errorf("Gagal mendapatkan data category")
			log.Println(err.Error())
			output = ResultUseCase{Error: err}
			return output
		}
		images, _ := resultCategory.Result.([]model.Image)
		products[i].Image = images
	}

	output = ResultUseCase{Result: products}
	return output
}

func (rp *ProductUsecaseImpl) UpdateProduct(productId string, param model.Product) ResultUseCase {
	output := ResultUseCase{}

	saveResult := rp.ProductRepository.UpdateProduct(productId, param)
	if saveResult.Error != nil {
		log.Println("Error mengupdate data :", saveResult.Error.Error())
		output = ResultUseCase{Error: saveResult.Error}
		return output
	}
	data, ok := saveResult.Result.(model.Product)
	if !ok {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		output = ResultUseCase{Error: err}
		return output
	}
	output = ResultUseCase{Result: data}
	return output
}

func (rp *ProductUsecaseImpl) DeleteProduct(productId string) ResultUseCase {
	output := ResultUseCase{}

	resultProduct := rp.ProductRepository.GetProductById(productId)
	if resultProduct.Error != nil {
		err := fmt.Errorf("Gagal mendapatkan data")
		log.Println(err.Error())
		output = ResultUseCase{Error: err}
		return output
	}
	products := resultProduct.Result.([]model.Product)
	for i := 0; i < len(products); i++ {
		id := strconv.Itoa(products[i].ID)
		resultCategory := rp.ProductRepository.GetImageOfProduct(id)
		if resultCategory.Error != nil {
			err := fmt.Errorf("Gagal mendapatkan data category")
			log.Println(err.Error())
			output = ResultUseCase{Error: err}
			return output
		}
		images, _ := resultCategory.Result.([]model.Image)
		products[i].Image = images
	}

	for i := 0; i < len(products); i++ {
		for j := 0; j < len(products[i].Image); j++ {
			id := strconv.Itoa(products[i].Image[j].ID)
			deleteResult := rp.ProductRepository.DeleteImageProduct(id)
			if deleteResult.Error != nil {
				log.Println("Error hapus data image :", deleteResult.Error.Error())
				output = ResultUseCase{Error: deleteResult.Error}
				return output
			}
		}
	}

	deleteResult := rp.ProductRepository.DeleteProduct(productId)
	if deleteResult.Error != nil {
		log.Println("Error hapus data :", deleteResult.Error.Error())
		output = ResultUseCase{Error: deleteResult.Error}
		return output
	}
	output = ResultUseCase{Result: productId}
	return output
}
