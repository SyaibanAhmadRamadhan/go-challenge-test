package schema

import (
	"github.com/gofiber/fiber/v2"
)

type CategoryProductRequest struct {
	Name string `json:"name"`
}

func (r *CategoryProductRequest) Validate() error {
	errBadRequest := map[string][]string{}

	if r.Name == "" {
		errBadRequest["name"] = append(errBadRequest["name"], Required)
	}
	name := MaxMinString(r.Name, 3, 100)
	if name != "" {
		errBadRequest["name"] = append(errBadRequest["name"], name)
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
