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
	filterStr, values, _ := repository.GenerateFilters(filters)
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM m_session %s)`, filterStr)

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
	filterStr, values, _ := repository.GenerateFilters(filters)
	query := fmt.Sprintf(`SELECT id, user_id, token, device, login_at, ip, %s FROM m_session 
                                                    %s`,
		repository.AuditToQuery(""), filterStr)

	tx, err := s.GetTx()
	if err != nil {
		return
	}

	session = &repository.Session{}

	err = tx.QueryRow(ctx, query, values...).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.Device,
		&session.LoginAt,
		&session.IP,
		&session.CreatedAt,
		&session.CreatedBy,
		&session.UpdatedAt,
		&session.UpdatedBy,
		&session.DeletedAt,
		&session.DeletedBy,
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			log.Warn().Msgf("failed query row | err: %v", err)
		}
		session = nil
		return
	}

	return
}
