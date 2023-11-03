package uow_repository

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func (u *UnitOfWorkRepositoryImpl) GetDB() (*pgxpool.Pool, error) {
	if u.db == nil {
		err := fmt.Errorf("no db Available")
		log.Warn().Msg(err.Error())
		return nil, err
	}

	return u.db, nil
}
