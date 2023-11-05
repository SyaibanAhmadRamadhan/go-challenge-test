package product_usecase

import (
	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

type ProductUsecaseImpl struct {
	productRepo         repository.ProductRepository
	categoryProductRepo repository.CategoryProductRepository
}

func NewProductUsecaseImpl(
	productRepo repository.ProductRepository,
	categoryProductRepo repository.CategoryProductRepository,
) usecase.ProductUsecase {
	return &ProductUsecaseImpl{
		productRepo:         productRepo,
		categoryProductRepo: categoryProductRepo,
	}
}
