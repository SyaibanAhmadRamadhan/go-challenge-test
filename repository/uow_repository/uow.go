package uow_repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"challenge-test-synapsis/repository"
)

type UnitOfWorkRepositoryImpl struct {
	tx pgx.Tx
	db *pgxpool.Pool
}

func NewUnitOfWorkRepositoryImpl(db *pgxpool.Pool) repository.UOWRepository {
	return &UnitOfWorkRepositoryImpl{
		db: db,
	}
}
