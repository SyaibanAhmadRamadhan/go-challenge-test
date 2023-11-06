package cart_repository

import (
	"challenge-test-synapsis/repository"
)

type CartRepositoryImpl struct {
	repository.UOWRepository
}

func NewCartRepositoryImpl(
	uow repository.UOWRepository,
) repository.CartRepository {
	return &CartRepositoryImpl{
		UOWRepository: uow,
	}
}
