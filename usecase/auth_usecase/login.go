package auth_usecase

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (a *AuthUsecaseImpl) Login(ctx context.Context, param *usecase.LoginParam) (res *usecase.AuthResult, err error) {
	filters := &[]repository.Filter{
		{
			Column:   "email",
			Value:    param.Email,
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
		user, err := a.userRepo.FindOne(ctx, filters)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				err = usecase.ErrInvalidEmailOrPass
			}
			return err
		}

		valid := usecase.CompareBcrypt(user.Password, param.Password)
		if !valid {
			log.Info().Msgf("password is not same")
			return usecase.ErrInvalidEmailOrPass
		}

		token, err := a.sessionManagementAuthCreate(ctx, &sessionManagementAuthCreate{
			UserID:   user.ID,
			Device:   param.Device,
			IP:       param.IP,
			TimeUnix: time.Now().Unix(),
			RoleID:   user.RoleID,
		})

		res = &usecase.AuthResult{
			ID:          user.ID,
			RoleID:      user.RoleID,
			Username:    user.Username,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Token:       token,
		}
		return err
	})

	return
}
