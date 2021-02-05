package usecase

import (
	"pretest-privyid/modules/v1/category/model"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */
type ResultUseCase struct {
	Result interface{}
	Error  error
}

type CategoryUsecase interface {
	CreateCategory(param model.Category) ResultUseCase
	GetAllCategories() ResultUseCase
	GetCategoryById(id string) ResultUseCase
	UpdateCategoryById(id string, param model.Category) ResultUseCase
	DeleteCategory(id string) ResultUseCase
}