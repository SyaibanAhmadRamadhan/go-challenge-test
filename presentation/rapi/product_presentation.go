package rapi

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/presentation/rapi/exception"
	"challenge-test-synapsis/presentation/rapi/schema"
	"challenge-test-synapsis/usecase"
)

func (p *Presenter) AddProduct(c *fiber.Ctx) error {
	req := new(schema.ProductRequest)

	if err := c.BodyParser(req); err != nil {
		return exception.Err(c, err)
	}

	err := req.Validate()
	if err != nil {
		return exception.Err(c, err)
	}

	product, err := p.ProductUsecase.Create(c.Context(), &usecase.ProductParam{
		CategoryProductID: req.CategoryProductID,
		Name:              req.Name,
		Stock:             req.Stock,
		Price:             req.Price,
		Description:       req.Description,
		CommonParam: usecase.CommonParam{
			UserID: c.Locals("userID").(string),
		},
	})

	if err != nil {
		if errors.Is(err, usecase.ErrCategoryProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "BAD REQUEST",
				Err: map[string][]string{
					"category_product_id": {
						err.Error(),
					},
				},
			}
		}
		if errors.Is(err, usecase.ErrProductNameIsExist) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "BAD REQUEST",
				Err: map[string][]string{
					"name": {
						err.Error(),
					},
				},
			}
		}

		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "successfully category product",
		Data: schema.ProductResponse{
			ID:                product.ID,
			CategoryProductID: product.CategoryProductID,
			Name:              product.Name,
			Price:             product.Price,
			Stock:             product.Stock,
			Description:       product.Description,
		},
		Err: nil,
	})
}

func (p *Presenter) UpdateProduct(c *fiber.Ctx) error {
	req := new(schema.ProductRequest)

	if err := c.BodyParser(req); err != nil {
		return exception.Err(c, err)
	}

	err := req.Validate()
	if err != nil {
		return exception.Err(c, err)
	}

	id := c.Params("id")
	_, err = helper.NewUlid(id)
	if err != nil {
		return exception.Err(c, &schema.ErrHttp{
			Code:    fiber.StatusNotFound,
			Message: "NOT FOUND",
			Err:     "product not found",
		})
	}

	product, err := p.ProductUsecase.Update(c.Context(), id, &usecase.ProductParam{
		CategoryProductID: req.CategoryProductID,
		Name:              req.Name,
		Stock:             req.Stock,
		Price:             req.Price,
		Description:       req.Description,
		CommonParam: usecase.CommonParam{
			UserID: c.Locals("userID").(string),
		},
	})

	if err != nil {
		if errors.Is(err, usecase.ErrCategoryProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "BAD REQUEST",
				Err: map[string][]string{
					"category_product_id": {
						err.Error(),
					},
				},
			}
		}
		if errors.Is(err, usecase.ErrProductNameIsExist) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "BAD REQUEST",
				Err: map[string][]string{
					"name": {
						err.Error(),
					},
				},
			}
		}
		if errors.Is(err, usecase.ErrProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusNotFound,
				Message: "NOT FOUND",
				Err:     "product not found",
			}
		}

		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "successfully updated product",
		Data: schema.ProductResponse{
			ID:                product.ID,
			CategoryProductID: product.CategoryProductID,
			Name:              product.Name,
			Price:             product.Price,
			Stock:             product.Stock,
			Description:       product.Description,
		},
		Err: nil,
	})
}

func (p *Presenter) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := helper.NewUlid(id)
	if err != nil {
		return exception.Err(c, &schema.ErrHttp{
			Code:    fiber.StatusNotFound,
			Message: "NOT FOUND",
			Err:     "product not found",
		})
	}

	err = p.ProductUsecase.Delete(c.Context(), id, &usecase.CommonParam{
		UserID: c.Locals("userID").(string),
	})

	if err != nil {
		if errors.Is(err, usecase.ErrProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusNotFound,
				Message: "NOT FOUND",
				Err:     "product not found",
			}
		}

		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "successfully deleted product",
		Data:    nil,
		Err:     nil,
	})
}

func (p *Presenter) GetProduct(c *fiber.Ctx) error {
	search := c.Query("search")
	page := c.Query("page")
	pageSize := c.Query("page-size")
	category := c.Query("category")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt > 10 {
		pageSizeInt = 10
	}

	products, paginate, err := p.ProductUsecase.GetAndSearch(c.Context(), search, &usecase.GetProductParam{
		CategoryProductID: category,
		PaginateParam: usecase.PaginateParam{
			Page:     pageInt,
			PageSize: pageSizeInt,
		},
	})

	if err != nil {
		return exception.Err(c, err)
	}

	var res []schema.ProductResponse

	for _, product := range *products {
		res = append(res, schema.ProductResponse{
			ID:                product.ID,
			CategoryProductID: product.CategoryProductID,
			Name:              product.Name,
			Price:             product.Price,
			Stock:             product.Stock,
			Description:       product.Description,
		})
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "data product",
		Data:    res,
		Err:     nil,
		Paginate: &schema.PaginateRes{
			CurrentPage: paginate.CurrentPage,
			Total:       paginate.Total,
			PageSize:    paginate.PageSize,
			PageTotal:   paginate.PageTotal,
		},
	})
}
