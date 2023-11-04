package session_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (s *SessionRepositoryImpl) CheckOne(ctx context.Context, filters *[]repository.Filter) (b bool, err error) {
	filterStr, values := repository.GenerateFilters(filters)
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM m_session WHERE %s AND deleted_at IS NULL)`, filterStr)

	tx, err := s.GetTx()
	if err != nil {
		return
	}

	err = tx.QueryRow(ctx, query, values...).Scan(&b)
	if err != nil {
		return
	}

	return
}

func (s *SessionRepositoryImpl) FindOne(ctx context.Context, filters *[]repository.Filter) (session *repository.Session, err error) {
	filterStr, values := repository.GenerateFilters(filters)
	query := fmt.Sprintf(`SELECT id, user_id, token, device, login_at, ip, %s FROM m_session 
                                                    WHERE %s AND deleted_at IS NULL`,
		repository.AuditToQuery(""), filterStr)

	tx, err := s.GetTx()
	if err != nil {
		return
	}

	var sessionScan repository.Session

	err = tx.QueryRow(ctx, query, values...).Scan(
		&sessionScan.ID,
		&sessionScan.UserID,
		&sessionScan.Token,
		&sessionScan.Device,
		&sessionScan.LoginAt,
		&sessionScan.IP,
		&sessionScan.CreatedAt,
		&sessionScan.CreatedBy,
		&sessionScan.UpdatedAt,
		&sessionScan.UpdatedBy,
		&sessionScan.DeletedAt,
		&sessionScan.DeletedBy,
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warn().Msgf("failed query row | err: %v", err)
		}
		return
	}

	session = &sessionScan
	return
}
