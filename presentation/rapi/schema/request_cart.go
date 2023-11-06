package schema

import (
	"github.com/gofiber/fiber/v2"

	"challenge-test-synapsis/helper"
)

type CartRequest struct {
	ProductID string `json:"product_id"`
	Total     int    `json:"total"`
}

func (r *CartRequest) Validate() error {
	errBadRequest := map[string][]string{}

	if r.ProductID == "" {
		errBadRequest["product_id"] = append(errBadRequest["product_id"], Required)
	}
	_, err := helper.NewUlid(r.ProductID)
	if err != nil {
		errBadRequest["product_id"] = append(errBadRequest["product_id"], "invalid product id")
	}

	total := MaxMinInt(r.Total, 1, 10)
	if total != "" {
		errBadRequest["total"] = append(errBadRequest["total"], total)
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

type CartRequestUpdate struct {
	Total int `json:"total"`
}

func (r *CartRequestUpdate) Validate() error {
	errBadRequest := map[string][]string{}

	total := MaxMinInt(r.Total, 1, 10)
	if total != "" {
		errBadRequest["total"] = append(errBadRequest["total"], total)
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
