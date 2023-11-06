package categoryProduct_usecase

import (
	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

type CategoryProductUsecaseImpl struct {
	categoryProductRepo repository.CategoryProductRepository
	productRepo         repository.ProductRepository
}

func NewCategoryProductUsecaseImpl(
	categoryProductRepo repository.CategoryProductRepository,
	productRepo repository.ProductRepository,
) usecase.CategoryProductUsecase {
	return &CategoryProductUsecaseImpl{
		categoryProductRepo: categoryProductRepo,
		productRepo:         productRepo,
	}
}
