package user_repository

import (
	"challenge-test-synapsis/repository"
)

type UserRepositoryImpl struct {
	repository.UOWRepository
}

func NewUserRepositoryImpl(uow repository.UOWRepository) repository.UserRepository {
	return &UserRepositoryImpl{
		UOWRepository: uow,
	}
}
