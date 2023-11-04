package user_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (u *UserRepositoryImpl) CheckOne(ctx context.Context, filters *[]repository.Filter) (b bool, err error) {
	filterStr, values := repository.GenerateFilters(filters)
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM m_user WHERE %s AND deleted_at IS NULL)", filterStr)

	tx, err := u.GetTx()
	if err != nil {
		return
	}

	err = tx.QueryRow(ctx, query, values...).Scan(&b)
	if err != nil {
		return
	}

	return
}

func (u *UserRepositoryImpl) FindOne(ctx context.Context, filters *[]repository.Filter) (user *repository.User, err error) {
	filterStr, values := repository.GenerateFilters(filters)
	query := fmt.Sprintf(`SELECT id, role_id, username, email, password, phone_number, %s 
									FROM m_user WHERE %s AND deleted_at IS NULL LIMIT 1`,
		repository.AuditToQuery(""), filterStr)

	tx, err := u.GetTx()
	if err != nil {
		return
	}

	var userScan repository.User

	err = tx.QueryRow(ctx, query, values...).Scan(
		&userScan.ID,
		&userScan.RoleID,
		&userScan.Username,
		&userScan.Email,
		&userScan.Password,
		&userScan.PhoneNumber,
		&userScan.CreatedAt,
		&userScan.CreatedBy,
		&userScan.UpdatedAt,
		&userScan.UpdatedBy,
		&userScan.DeletedAt,
		&userScan.DeletedBy,
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warn().Msgf("failed query row | err: %v", err)
		}
		return
	}

	user = &userScan
	return
}
