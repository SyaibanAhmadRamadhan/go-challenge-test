package conf

import (
	"os"
	"time"
)

type JwtConf struct {
	ATkey         string
	ATexp         time.Duration
	RTkey         string
	RTexp         time.Duration
	RememberMeExp time.Duration
}

func EnvJwtConf() JwtConf {
	atKey := os.Getenv("AUTH_JWT_TOKEN_HS_ACCESS_TOKEN_KEY")
	if atKey == "" {
		panic("AUTH_JWT_TOKEN_HS_ACCESS_TOKEN_KEY is nil")
	}
	atExp := os.Getenv("AUTH_JWT_TOKEN_HS_ACCESS_TOKEN_EXPIRED")
	if atExp == "" {
		panic("AUTH_JWT_TOKEN_HS_ACCESS_TOKEN_EXPIRED is nil")
	}
	rtKey := os.Getenv("AUTH_JWT_TOKEN_HS_REFRESH_TOKEN_KEY")
	if rtKey == "" {
		panic("AUTH_JWT_TOKEN_HS_REFRESH_TOKEN_KEY is nil")
	}
	rtExp := os.Getenv("AUTH_JWT_TOKEN_HS_REFRESH_TOKEN_EXPIRED")
	if rtExp == "" {
		panic("AUTH_JWT_TOKEN_HS_REFRESH_TOKEN_EXPIRED is nil")
	}
	rememberMeExp := os.Getenv("AUTH_JWT_TOKEN_REMEMBER_ME_EXPIRED")
	if rememberMeExp == "" {
		panic("AUTH_JWT_TOKEN_REMEMBER_ME_EXPIRED is nil")
	}

	atExpTime, err := time.ParseDuration(atExp)
	if err != nil {
		panic(err)
	}
	rtExpTime, err := time.ParseDuration(rtExp)
	if err != nil {
		panic(err)
	}
	rememberMeExpTime, err := time.ParseDuration(rememberMeExp)
	if err != nil {
		panic(err)
	}

	return JwtConf{
		ATkey:         atKey,
		ATexp:         atExpTime,
		RTkey:         rtKey,
		RTexp:         rtExpTime,
		RememberMeExp: rememberMeExpTime,
	}
}
