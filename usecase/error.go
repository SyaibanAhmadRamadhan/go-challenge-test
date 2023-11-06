package usecase

import (
	"errors"
)

var ErrJwtSigningMethodInvalid = errors.New("invalid signing method jwt")
var ErrEmailIsRegistered = errors.New("email is registered")
var ErrInvalidEmailOrPass = errors.New("invalid email or password")
var ErrInvalidToken = errors.New("invalid token")
var ErrCategoryProductNameIsExist = errors.New("category product name is available")
var ErrCategoryProductNotFound = errors.New("category product not found")
var ErrCategoryProductHaveProduct = errors.New("category product have any product, cannot deleted category product")
var ErrProductNameIsExist = errors.New("product name is available")
var ErrProductNotFound = errors.New("product not found")
var ErrItemCartNotFound = errors.New("your item cart product not found")
