package session_repository

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
)

func (s *SessionRepositoryImpl) Create(ctx context.Context, session *repository.Session) (err error) {
	query := `INSERT INTO m_session (id, user_id, token, device, login_at, ip, created_at, created_by, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	tx, err := s.GetTx()
	if err != nil {
		return
	}

	_, err = tx.Exec(ctx, query,
		session.ID,
		session.UserID,
		session.Token,
		session.Device,
		session.LoginAt,
		session.IP,
		session.CreatedAt,
		session.CreatedBy,
		session.UpdatedAt,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	return
}

func (s *SessionRepositoryImpl) Update(ctx context.Context, session *repository.Session) (err error) {
	query := `UPDATE m_session SET token=$1, updated_at=$2, updated_by=$3 WHERE id=$4 AND user_id=$5 AND deleted_at IS NULL`

	tx, err := s.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		session.Token,
		session.UpdatedAt,
		session.UserID,
		session.ID,
		session.UserID,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	if res.RowsAffected() == 0 {
		log.Info().Msgf("updated does not meet the conditions")
	}

	return
}

func (s *SessionRepositoryImpl) Delete(ctx context.Context, id int, userID string) (err error) {
	query := `UPDATE m_session SET deleted_at=$1, deleted_by=$2 WHERE id=$3 AND user_id=$4 AND deleted_at IS NULL`

	tx, err := s.GetTx()
	if err != nil {
		return
	}

	res, err := tx.Exec(ctx, query,
		time.Now().Unix(),
		userID,
		id,
		userID,
	)
	if err != nil {
		log.Warn().Msgf("failed exec command | err: %v", err)
		return
	}

	if res.RowsAffected() == 0 {
		log.Info().Msgf("updated does not meet the conditions")
	}

	return
}
