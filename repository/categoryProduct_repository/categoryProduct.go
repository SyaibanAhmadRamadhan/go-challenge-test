package categoryProduct_repository

import (
	"challenge-test-synapsis/repository"
)

type CategoryProductRepositoryImpl struct {
	repository.UOWRepository
}

func NewCategoryProductRepositoryImpl(uow repository.UOWRepository) repository.CategoryProductRepository {
	return &CategoryProductRepositoryImpl{
		UOWRepository: uow,
	}
}
