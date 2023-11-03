package user_repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (u *UserRepositoryImpl) Create(ctx context.Context, user *repository.User) (err error) {
	query := `INSERT INTO m_user (id, role_id, username, email, password, phone_number, created_at, created_by, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	tx, err := u.GetTx()
	if err != nil {
		return
	}

	_, err = tx.Exec(ctx, query,
		user.ID,
		user.Role.ID,
		user.Username,
		user.Email,
		user.Password,
		user.PhoneNumber,
		user.CreatedAt,
		user.CreatedBy,
		user.UpdatedAt,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	return
}

func (u *UserRepositoryImpl) Update(ctx context.Context, user *repository.User) (err error) {
	query := `UPDATE m_user SET role_id=$1, username=$2, email=$3, password=$4, phone_number=$5, updated_at=$6, updated_by=$7
				WHERE id = $8 AND deleted_at IS NULL`

	tx, err := u.GetTx()
	if err != nil {
		return
	}

	_, err = tx.Exec(ctx, query,
		user.Role.ID,
		user.Username,
		user.Email,
		user.Password,
		user.PhoneNumber,
		user.UpdatedAt,
		user.UpdatedBy,
		user.ID,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	return
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, id string) (err error) {
	query := `UPDATE m_user SET deleted_at=$1
				WHERE id = $2 AND deleted_at IS NULL`

	tx, err := u.GetTx()
	if err != nil {
		return
	}

	_, err = tx.Exec(ctx, query,
		time.Now().Unix(),
		id,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	return
}
