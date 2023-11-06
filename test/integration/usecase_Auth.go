package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"challenge-test-synapsis/usecase"
)

var token string
var id string
var device string

func AuthUsecaseImplRegister(t *testing.T) {
	register := &usecase.RegisterParam{
		RoleID:      2,
		Username:    "rama",
		Email:       "test@gmail.com",
		Password:    "rama123",
		PhoneNumber: "088295007524",
		RememberMe:  true,
		CommonParam: usecase.CommonParam{
			Device: "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36",
			IP:     "127.0.0.1",
		},
	}

	res, err := AuthUsecase.Register(context.Background(), register)
	assert.NoError(t, err)
	assert.Equal(t, 2, res.RoleID)
	t.Log(res)
}

func AuthUsecaseImplLogin(t *testing.T) {
	login := &usecase.LoginParam{
		Email:    "test@gmail.com",
		Password: "rama123",
		CommonParam: usecase.CommonParam{
			Device: "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36",
			IP:     "127.0.0.1",
		},
	}

	res, err := AuthUsecase.Login(context.Background(), login)
	assert.NoError(t, err)
	assert.Equal(t, 2, res.RoleID)
	token = res.Token
	id = res.ID
	device = login.Device
	t.Log(res)
}

func AuthUsecaseImplOtorisasi(t *testing.T) {
	t.Run("token_in_redis", func(t *testing.T) {
		res, err := AuthUsecase.Otorisasi(context.Background(), token, &usecase.CommonParam{
			UserID: id,
			Device: device,
		})
		assert.NoError(t, err)
		t.Log(res)
	})

	newToken := ""
	t.Run("generate_a_new_token", func(t *testing.T) {
		res, err := AuthUsecase.Otorisasi(context.Background(), sessionExp.Token, &usecase.CommonParam{
			UserID: sessionExp.UserID,
			Device: device,
		})
		newToken = res.Token
		assert.NoError(t, err)
		t.Log(res)
	})

	t.Run("validate_new_token", func(t *testing.T) {
		res, err := AuthUsecase.Otorisasi(context.Background(), newToken, &usecase.CommonParam{
			UserID: sessionExp.UserID,
			Device: device,
		})
		assert.NoError(t, err)
		t.Log(res)
	})

	t.Run("token_invalid", func(t *testing.T) {
		token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTkxNzExOTksInN1YiI6IjAxSEVGNlM2RTdTOUQzNFE4In0.6kDtcMj6nBZW-5rCrRzg1f3l2NdPJogKvOV2UGzI8js"
		res, err := AuthUsecase.Otorisasi(context.Background(), token, &usecase.CommonParam{
			UserID: id,
			Device: device,
		})
		assert.ErrorIs(t, err, usecase.ErrInvalidToken)
		t.Log(res)
	})
}

func AuthUsecaseImplRegisterError(t *testing.T) {
	register := &usecase.RegisterParam{
		RoleID:      1,
		Username:    "rama",
		Email:       "test@gmail.com",
		Password:    "rama123",
		PhoneNumber: "088295007524",
		RememberMe:  true,
		CommonParam: usecase.CommonParam{
			Device: "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36",
			IP:     "127.0.0.1",
		},
	}

	t.Run("email_is_registered", func(t *testing.T) {
		res, err := AuthUsecase.Register(context.Background(), register)
		assert.ErrorIs(t, err, usecase.ErrEmailIsRegistered)
		assert.Nil(t, res)
	})
}

func AuthUsecaseImplLoginError(t *testing.T) {
	login := &usecase.LoginParam{
		Email:    "test@gmail.com",
		Password: "rama123",
		CommonParam: usecase.CommonParam{
			Device: "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36",
			IP:     "127.0.0.1",
		},
	}

	t.Run("invalid_email", func(t *testing.T) {
		login.Email = "asal"
		res, err := AuthUsecase.Login(context.Background(), login)
		assert.ErrorIs(t, err, usecase.ErrInvalidEmailOrPass)
		assert.Nil(t, res)
	})

	t.Run("invalid_password", func(t *testing.T) {
		login.Password = "asal"
		res, err := AuthUsecase.Login(context.Background(), login)
		assert.ErrorIs(t, err, usecase.ErrInvalidEmailOrPass)
		assert.Nil(t, res)
	})
}
