package usecase

import (
	"pretest-privyid/modules/v1/product/model"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */
type ResultUseCase struct {
	Result interface{}
	Error  error
}

type ProductUsecase interface {
	CreateProduct(param model.Product) ResultUseCase
	UploadImage(productId string, param model.Image) ResultUseCase
	GetAllProduct() ResultUseCase
	GetProductById(productId string) ResultUseCase
	UpdateProduct(productId string, param model.Product) ResultUseCase
	DeleteProduct(productId string) ResultUseCase
}