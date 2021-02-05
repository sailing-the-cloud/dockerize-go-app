package repository

import "pretest-privyid/modules/v1/product/model"

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type ResultRepository struct {
	Result interface{}
	Error  error
}

type ProductRepository interface {
	CreateProduct(param model.Product) ResultRepository
	UploadImage(productId string, param model.Image) ResultRepository
	GetAllProduct() ResultRepository
	GetCategoryOfProduct(productId string) ResultRepository
	GetImageOfProduct(productId string) ResultRepository
	GetProductById(productId string) ResultRepository
	UpdateProduct(productId string, param model.Product) ResultRepository
	DeleteProduct(productId string) ResultRepository
	DeleteImageProduct(imageId string) ResultRepository
}