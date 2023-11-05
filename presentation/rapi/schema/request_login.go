package schema

import (
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	errBadRequest := map[string][]string{}

	if r.Email == "" {
		errBadRequest["email"] = append(errBadRequest["email"], Required)
	}
	email := MaxMinString(r.Email, 3, 55)
	if email != "" {
		errBadRequest["email"] = append(errBadRequest["email"], email)
	}

	if r.Password == "" {
		errBadRequest["password"] = append(errBadRequest["password"], Required)
	}
	password := MaxMinString(r.Password, 6, 55)
	if password != "" {
		errBadRequest["password"] = append(errBadRequest["password"], password)
	}

	if len(errBadRequest) != 0 {
		return &ErrorHttp{
			Code:    fiber.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    nil,
			Err:     errBadRequest,
		}
	}
	return nil
}
