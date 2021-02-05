package repository

import "pretest-privyid/modules/v1/category/model"

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type ResultRepository struct {
	Result interface{}
	Error  error
}

type CategoryRepository interface {
	CreateCategory(param model.Category) ResultRepository
	GetAllCategories() ResultRepository
	GetCategoryById(id string) ResultRepository
	UpdateCategoryById(id string, param model.Category) ResultRepository
	DeleteCategory(id string) ResultRepository
}