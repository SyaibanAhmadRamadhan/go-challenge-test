package auth_usecase

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"challenge-test-synapsis/conf"
	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/repository"
	"challenge-test-synapsis/usecase"
)

func (a *AuthUsecaseImpl) Register(ctx context.Context, param *usecase.RegisterParam) (res *usecase.AuthResult, err error) {
	passwordHash, err := usecase.HashBcrypt(param.Password)
	if err != nil {
		return
	}
	id, _ := helper.NewUlid(helper.EmptyString)
	timeUnix := time.Now().Unix()

	admin := conf.EnvAdmin()
	if admin[param.Email] {
		param.RoleID = 2
	}

	filterCheckMailAndPhoneNumber := &[]repository.Filter{
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
		exist, err := a.userRepo.CheckOne(ctx, filterCheckMailAndPhoneNumber)
		if exist {
			err = usecase.ErrEmailIsRegistered
			return err
		}

		err = a.userRepo.Create(ctx, &repository.User{
			ID:          id,
			RoleID:      param.RoleID,
			Username:    param.Username,
			Email:       param.Email,
			Password:    passwordHash,
			PhoneNumber: param.PhoneNumber,
			Audit: repository.Audit{
				CreatedAt: timeUnix,
				CreatedBy: id,
				UpdatedAt: timeUnix,
			},
		})
		if err != nil {
			return err
		}

		token, err := a.sessionManagementAuthCreate(ctx, &sessionManagementAuthCreate{
			UserID:   id,
			Device:   param.Device,
			IP:       param.IP,
			TimeUnix: timeUnix,
			RoleID:   param.RoleID,
		})

		res = &usecase.AuthResult{
			ID:          id,
			RoleID:      param.RoleID,
			Username:    param.Username,
			Email:       param.Email,
			PhoneNumber: param.PhoneNumber,
			Token:       token,
		}
		return err
	})

	return
}
