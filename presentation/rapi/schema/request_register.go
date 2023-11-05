package schema

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RePassword  string `json:"re_password"`
	PhoneNumber string `json:"phone_number"`
}

func (r *RegisterRequest) Validate() error {
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
	if r.Password != r.RePassword {
		errBadRequest["password"] = append(errBadRequest["password"], PasswordAndRePassword)
		errBadRequest["re_password"] = append(errBadRequest["re_password"], PasswordAndRePassword)
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
