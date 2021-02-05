package main

import (
	"pretest-privyid/config"
	CategoryPresenterPackage "pretest-privyid/modules/v1/category/presenter"
	CategoryRepoPackage "pretest-privyid/modules/v1/category/repository"
	CategoryUseCasePackage "pretest-privyid/modules/v1/category/usecase"

	ProductPresenterPackage "pretest-privyid/modules/v1/product/presenter"
	ProductRepoPackage "pretest-privyid/modules/v1/product/repository"
	ProductUseCasePackage "pretest-privyid/modules/v1/product/usecase"
)

// Service structure
type Service struct {
	CategoryHandler *CategoryPresenterPackage.HTTPCategoryHandler
	CategoryUsecase CategoryUseCasePackage.CategoryUsecase

	ProductHandler *ProductPresenterPackage.HTTPProductHandler
	ProductUsecase ProductUseCasePackage.ProductUsecase
}

// MakeHandler function for initializing service
func MakeHandler() *Service {

	dbConnection := config.ConnectDB()
	CategoryRepo := CategoryRepoPackage.NewWorkCategoryRepoPostgres(dbConnection)
	CategoryUsecase := CategoryUseCasePackage.NewWorkCategoryUseCase(CategoryRepo)
	CategoryHandler := CategoryPresenterPackage.NewHTTPHandler(CategoryUsecase)

	ProductRepo := ProductRepoPackage.NewProductRepoPostgres(dbConnection)
	ProductUsecase := ProductUseCasePackage.NewProductUseCase(ProductRepo)
	ProductHandler := ProductPresenterPackage.NewHTTPHandler(ProductUsecase)

	return &Service{
		CategoryHandler: CategoryHandler,
		CategoryUsecase: CategoryUsecase,
		ProductHandler: ProductHandler,
		ProductUsecase: ProductUsecase,
	}
}
