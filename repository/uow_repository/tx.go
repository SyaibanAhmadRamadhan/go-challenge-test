package uow_repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (u *UnitOfWorkRepositoryImpl) StartTx(ctx context.Context, opts pgx.TxOptions, fn func() error) error {
	if u.db == nil {
		err := fmt.Errorf("no Connection Database Available")
		log.Warn().Msg(err.Error())
		return err
	}

	tx, err := u.db.BeginTx(ctx, opts)
	if err != nil {
		log.Warn().Msgf("failed start begin tx | err : %v", err)
		return err
	}
	u.tx = tx

	err = fn()
	if err != nil {
		if errRollback := tx.Rollback(ctx); errRollback != nil {
			log.Warn().Msgf("failed rollback data | err : %v", errRollback)
			return errRollback
		}

		log.Info().Msgf("have error, rollback data | err: %v", err)
		return err
	}

	if errCommit := tx.Commit(ctx); errCommit != nil {
		log.Warn().Msgf("failed commit data | %err : %v", errCommit)
		return errCommit
	}

	log.Info().Msgf("sucessfully, commit data")
	return nil
}

func (u *UnitOfWorkRepositoryImpl) GetTx() (pgx.Tx, error) {
	if u.tx == nil {
		log.Warn().Msg(repository.ErrTxIsNil.Error())
		return nil, repository.ErrTxIsNil
	}

	return u.tx, nil
}
