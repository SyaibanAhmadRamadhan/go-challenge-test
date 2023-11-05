package auth_usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/conf"
	"challenge-test-synapsis/usecase"
)

func (a *AuthUsecaseImpl) Otorisasi(ctx context.Context, param *usecase.OtorisasiParam) (res *usecase.OtorisasiResult, err error) {
	confJwt := conf.EnvJwtConf()
	claims, err := usecase.ClaimJwtHS256(param.Token, confJwt.ATkey)
	if err != nil {
		if !errors.Is(err, jwt.ErrTokenExpired) {
			return res, usecase.ErrInvalidToken
		}

		token, role, err := a.sessionManagementAuthUpdate(ctx, &sessionManagementAuthUpdate{
			UserID: param.UserID,
			Device: param.Device,
			Token:  param.Token,
		})

		res = &usecase.OtorisasiResult{
			RoleID: role,
			Token:  token,
		}

		return res, err
	}

	sessionID, ok := claims["sub"].(string)
	if !ok {
		log.Warn().Msgf("cannot assetion claims sub to string | data : %v", claims["sub"])
		return res, usecase.ErrInvalidToken
	}

	session, err := a.redisRepo.Get(ctx, fmt.Sprintf("%s:%s", sessionID, param.UserID))
	if err != nil {
		return res, usecase.ErrInvalidToken
	}

	sessionArr := strings.Split(session, "|")
	if len(sessionArr) != 3 {
		return res, usecase.ErrInvalidToken
	}

	if sessionArr[1] != param.Device {
		log.Info().Msgf("device is not the same")
		return res, usecase.ErrInvalidToken
	}

	roleInt, _ := strconv.Atoi(sessionArr[2])
	res = &usecase.OtorisasiResult{
		Token:  param.Token,
		RoleID: roleInt,
	}
	return
}
