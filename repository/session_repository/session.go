package session_repository

import (
	"challenge-test-synapsis/repository"
)

type SessionRepositoryImpl struct {
	repository.UOWRepository
}

func NewSessionRepositoryImpl(uow repository.UOWRepository) repository.SessionRepository {
	return &SessionRepositoryImpl{
		UOWRepository: uow,
	}
}
