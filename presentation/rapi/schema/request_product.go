package schema

import (
	"github.com/gofiber/fiber/v2"

	"challenge-test-synapsis/helper"
)

type ProductRequest struct {
	Name              string `json:"name"`
	Price             int    `json:"price"`
	Stock             int    `json:"stock"`
	Description       string `json:"description"`
	CategoryProductID string `json:"category_product_id"`
}

func (r *ProductRequest) Validate() error {
	errBadRequest := map[string][]string{}

	if r.CategoryProductID == "" {
		errBadRequest["category_product_id"] = append(errBadRequest["category_product_id"], Required)
	}
	_, err := helper.NewUlid(r.CategoryProductID)
	if err != nil {
		errBadRequest["category_product_id"] = append(errBadRequest["category_product_id"], "invalid category product id")
	}

	if r.Name == "" {
		errBadRequest["name"] = append(errBadRequest["name"], Required)
	}
	name := MaxMinString(r.Name, 3, 100)
	if name != "" {
		errBadRequest["name"] = append(errBadRequest["name"], name)
	}

	if r.Price == 0 {
		errBadRequest["price"] = append(errBadRequest["price"], Required)
	}
	price := MaxMinInt(r.Price, 1000, 100000000)
	if price != "" {
		errBadRequest["price"] = append(errBadRequest["price"], price)
	}

	stock := MaxMinInt(r.Stock, 0, 10000)
	if stock != "" {
		errBadRequest["stock"] = append(errBadRequest["stock"], stock)
	}

	if r.Description == "" {
		errBadRequest["description"] = append(errBadRequest["description"], Required)
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
