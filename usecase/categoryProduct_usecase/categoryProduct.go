package categoryProduct_usecase

import (
	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

type CategoryProductUsecaseImpl struct {
	categoryProductRepo repository.CategoryProductRepository
}

func NewCategoryProductUsecaseImpl(
	categoryProductRepo repository.CategoryProductRepository,
) usecase.CategoryProductUsecase {
	return &CategoryProductUsecaseImpl{
		categoryProductRepo: categoryProductRepo,
	}
}
