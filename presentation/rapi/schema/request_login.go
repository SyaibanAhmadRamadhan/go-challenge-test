package schema

import (
	"regexp"

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
	email := MaxMinString(r.Email, 12, 55)
	if email != "" {
		errBadRequest["email"] = append(errBadRequest["email"], email)
	}

	match, err := regexp.MatchString(`^([A-Za-z.]|[0-9])+@gmail.com$`, r.Email)
	if err != nil || !match {
		errBadRequest["email"] = append(errBadRequest["email"], EmailMsg)
	}

	if r.Password == "" {
		errBadRequest["password"] = append(errBadRequest["password"], Required)
	}
	password := MaxMinString(r.Password, 6, 55)
	if password != "" {
		errBadRequest["password"] = append(errBadRequest["password"], password)
	}

	if len(errBadRequest) != 0 {
		return &ErrHttp{
			Code:    fiber.StatusBadRequest,
			Message: "BAD REQUEST",
			Err:     errBadRequest,
		}
	}
	return nil
}
