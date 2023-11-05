package schema

import (
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Username    string `json:"username,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

func (r *RegisterRequest) Validate() error {
	errBadRequest := map[string][]string{}

	if r.Email == "" {
		errBadRequest["email"] = append(errBadRequest["email"], Required)
	}
	email := MaxMinString(r.Email, 3, 55)
	if email != "" {
		errBadRequest["email"] = append(errBadRequest["email"], email)
	}

	if r.Username == "" {
		errBadRequest["username"] = append(errBadRequest["username"], Required)
	}
	username := MaxMinString(r.Username, 3, 50)
	if username != "" {
		errBadRequest["username"] = append(errBadRequest["username"], username)
	}

	if r.PhoneNumber == "" {
		errBadRequest["phone_number"] = append(errBadRequest["phone_number"], Required)
	}
	phoneNumber := MaxMinString(r.PhoneNumber, 3, 15)
	if phoneNumber != "" {
		errBadRequest["phone_number"] = append(errBadRequest["phone_number"], phoneNumber)
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
