package product_repository

import (
	"challenge-test-synapsis/repository"
)

type ProductRepositoryImpl struct {
	repository.UOWRepository
}

func NewProductRepositoryImpl(uow repository.UOWRepository) repository.ProductRepository {
	return &ProductRepositoryImpl{
		UOWRepository: uow,
	}
}
