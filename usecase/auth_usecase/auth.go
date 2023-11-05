package auth_usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/conf"
	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

type AuthUsecaseImpl struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
	redisRepo   repository.RedisRepository
}

func NewAuthUsecaseImpl(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	redisRepo repository.RedisRepository,
) usecase.AuthUsecase {
	return &AuthUsecaseImpl{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		redisRepo:   redisRepo,
	}
}

type sessionManagementAuthCreate struct {
	UserID   string
	Device   string
	IP       string
	TimeUnix int64
	RoleID   int
}

type sessionManagementAuthUpdate struct {
	UserID string
	Device string
	Token  string
}

func (a *AuthUsecaseImpl) sessionManagementAuthCreate(ctx context.Context, session *sessionManagementAuthCreate) (token string, err error) {
	sessionID, _ := helper.NewUlid(helper.EmptyString)

	jwtModel := &usecase.Jwt{
		Conf: conf.EnvJwtConf(),
	}
	atModel := jwtModel.ATDefault(sessionID)
	at, err := usecase.GenerateJwtHS256(atModel)
	if err != nil {
		return
	}

	err = a.sessionRepo.Create(ctx, &repository.Session{
		ID:      sessionID,
		UserID:  session.UserID,
		Token:   at,
		Device:  session.Device,
		LoginAt: session.TimeUnix,
		IP:      session.IP,
		Audit: repository.Audit{
			CreatedAt: session.TimeUnix,
			CreatedBy: session.UserID,
			UpdatedAt: session.TimeUnix,
		},
	})
	if err != nil {
		return
	}

	err = a.redisRepo.Set(ctx, repository.RedisSet{
		Key:   fmt.Sprintf("%s:%s", sessionID, session.UserID),
		Value: fmt.Sprintf("%s|%s|%d", session.IP, session.Device, session.RoleID),
		Exp:   atModel.Exp,
	})

	return at, err
}

func (a *AuthUsecaseImpl) sessionManagementAuthUpdate(ctx context.Context, param *sessionManagementAuthUpdate) (at string, roleID int, err error) {
	filters := &[]repository.Filter{
		{
			Column:   "token",
			Value:    param.Token,
			Operator: repository.Equality,
		},
		{
			Column:   "user_id",
			Value:    param.UserID,
			Operator: repository.Equality,
		},
		{
			Column:   "deleted_at",
			Operator: repository.IsNULL,
		},
	}

	err = a.userRepo.StartTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	}, func() error {
		session, err := a.sessionRepo.FindOne(ctx, filters)
		if err != nil {
			return err
		}

		if session.Device != param.Device {
			return usecase.ErrInvalidToken
		}

		user, err := a.userRepo.FindOne(ctx, &[]repository.Filter{
			{
				Column:   "id",
				Value:    session.UserID,
				Operator: repository.Equality,
			},
			{
				Column:   "deleted_at",
				Operator: repository.IsNULL,
			},
		})
		if err != nil {
			return err
		}

		jwtModel := &usecase.Jwt{
			Conf: conf.EnvJwtConf(),
		}
		atModel := jwtModel.ATDefault(session.ID)
		at, err = usecase.GenerateJwtHS256(atModel)
		if err != nil {
			return err
		}

		err = a.sessionRepo.Update(ctx, &repository.Session{
			ID:     session.ID,
			UserID: session.UserID,
			Token:  at,
			Audit: repository.Audit{
				UpdatedAt: time.Now().Unix(),
				UpdatedBy: helper.NewNullString(session.UserID),
			},
		})
		if err != nil {
			return err
		}

		err = a.redisRepo.Set(ctx, repository.RedisSet{
			Key:   fmt.Sprintf("%s:%s", session.ID, session.UserID),
			Value: fmt.Sprintf("%s|%s|%d", session.IP, session.Device, user.RoleID),
			Exp:   atModel.Exp,
		})

		roleID = user.RoleID
		return err
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = usecase.ErrInvalidToken
		}
	}

	return
}
