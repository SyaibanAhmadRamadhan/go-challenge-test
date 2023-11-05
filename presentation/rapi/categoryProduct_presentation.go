package rapi

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/presentation/rapi/exception"
	"challenge-test-synapsis/presentation/rapi/schema"
	"challenge-test-synapsis/usecase"
)

func (p *Presenter) AddCategoryProduct(c *fiber.Ctx) error {
	req := new(schema.CategoryProductRequest)

	if err := c.BodyParser(req); err != nil {
		return exception.Err(c, err)
	}

	err := req.Validate()
	if err != nil {
		return exception.Err(c, err)
	}

	categoryProduct, err := p.CategoryProductUsecase.Create(c.Context(), &usecase.CategoryProductParam{
		Name: req.Name,
		CommonParam: usecase.CommonParam{
			UserID: c.Locals("userID").(string),
		},
	})

	if err != nil {
		if errors.Is(err, usecase.ErrCategoryProductNameIsExist) {
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
		Message: "successfully created category product",
		Data: schema.CategoryProductResponse{
			ID:   categoryProduct.ID,
			Name: categoryProduct.Name,
		},
		Err: nil,
	})
}

func (p *Presenter) UpdateCategoryProduct(c *fiber.Ctx) error {
	req := new(schema.CategoryProductRequest)

	if err := c.BodyParser(req); err != nil {
		return exception.Err(c, err)
	}

	err := req.Validate()
	if err != nil {
		return exception.Err(c, err)
	}

	id := c.Params("id")
	categoryProduct, err := p.CategoryProductUsecase.Update(c.Context(), id, &usecase.CategoryProductParam{
		Name: req.Name,
		CommonParam: usecase.CommonParam{
			UserID: c.Locals("userID").(string),
		},
	})

	if err != nil {
		if errors.Is(err, usecase.ErrCategoryProductNameIsExist) {
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
		if errors.Is(err, usecase.ErrCategoryProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusNotFound,
				Message: "NOT FOUND",
				Err:     err.Error(),
			}
		}

		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "successfully updated category product",
		Data: schema.CategoryProductResponse{
			ID:   categoryProduct.ID,
			Name: categoryProduct.Name,
		},
		Err: nil,
	})
}

func (p *Presenter) DeleteCategoryProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	err := p.CategoryProductUsecase.Delete(c.Context(), id, &usecase.CommonParam{
		UserID: c.Locals("userID").(string),
	})

	if err != nil {
		if errors.Is(err, usecase.ErrCategoryProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusNotFound,
				Message: "NOT FOUND",
				Err:     err.Error(),
			}
		}

		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "successfully deleted category product",
		Data:    nil,
		Err:     nil,
	})
}

func (p *Presenter) GetCategoryProduct(c *fiber.Ctx) error {
	search := c.Query("search")
	page := c.Query("page")
	pageSize := c.Query("page-size")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt > 10 {
		pageSizeInt = 10
	}

	categoryProducts, paginate, err := p.CategoryProductUsecase.GetAndSearch(c.Context(), search, &usecase.PaginateParam{
		Page:     pageInt,
		PageSize: pageSizeInt,
	})

	if err != nil {
		return exception.Err(c, err)
	}

	log.Info().Msgf("%v", categoryProducts)
	var res []schema.CategoryProductResponse

	for _, categoryProduct := range *categoryProducts {
		res = append(res, schema.CategoryProductResponse{
			ID:   categoryProduct.ID,
			Name: categoryProduct.Name,
		})
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "successfully deleted category product",
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
