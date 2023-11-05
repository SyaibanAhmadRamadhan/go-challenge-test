package usecase

import (
	"context"
)

type RegisterParam struct {
	RoleID      int
	Username    string
	Email       string
	Password    string
	PhoneNumber string
	RememberMe  bool
	CommonParam
}

type LoginParam struct {
	Email    string
	Password string
	CommonParam
}

type AuthResult struct {
	ID          string
	RoleID      int
	Username    string
	Email       string
	PhoneNumber string
	Token       string
}

type OtorisasiResult struct {
	RoleID int
	Token  string
}

type OtorisasiParam struct {
	UserID string
	Token  string
	Device string
}

type AuthUsecase interface {
	Register(ctx context.Context, param *RegisterParam) (res *AuthResult, err error)
	Login(ctx context.Context, param *LoginParam) (res *AuthResult, err error)
	Otorisasi(ctx context.Context, param *OtorisasiParam) (res *OtorisasiResult, err error)
}
