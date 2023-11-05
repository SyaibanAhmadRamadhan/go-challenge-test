package usecase

import (
	"errors"
)

var ErrJwtSigningMethodInvalid = errors.New("invalid signing method jwt")
var ErrEmailIsRegistered = errors.New("email is registered")
var ErrInvalidEmailOrPass = errors.New("invalid email or password")
var ErrInvalidToken = errors.New("invalid token")
