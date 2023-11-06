package rapi

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"challenge-test-synapsis/presentation/rapi/exception"
	"challenge-test-synapsis/presentation/rapi/schema"
	"challenge-test-synapsis/usecase"
)

func (p *Presenter) AddProductCart(c *fiber.Ctx) error {
	req := new(schema.CartRequest)

	if err := c.BodyParser(req); err != nil {
		return exception.Err(c, err)
	}

	err := req.Validate()
	if err != nil {
		return exception.Err(c, err)
	}

	cart, err := p.CartUsecase.AddItemCart(c.Context(), &usecase.ItemCartParam{
		ProductID: req.ProductID,
		Total:     req.Total,
		CommonParam: usecase.CommonParam{
			UserID: c.Locals("userID").(string),
		},
	})

	if err != nil {
		if errors.Is(err, usecase.ErrProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "BAD REQUEST",
				Err: map[string][]string{
					"product_id": {
						err.Error(),
					},
				},
			}
		}
		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "created list cart successfully",
		Data:    cart,
		Err:     nil,
	})
}

func (p *Presenter) UpdateProductCart(c *fiber.Ctx) error {
	req := new(schema.CartRequestUpdate)

	if err := c.BodyParser(req); err != nil {
		return exception.Err(c, err)
	}

	err := req.Validate()
	if err != nil {
		return exception.Err(c, err)
	}

	id := c.Params("id")
	// _, err = helper.NewUlid(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return exception.Err(c, &schema.ErrHttp{
			Code:    fiber.StatusNotFound,
			Message: "NOT FOUND",
			Err:     "cart not found",
		})
	}

	cart, err := p.CartUsecase.UpdateItemCart(c.Context(), int64(idInt), &usecase.ItemCartParam{
		Total: req.Total,
		CommonParam: usecase.CommonParam{
			UserID: c.Locals("userID").(string),
		},
	})

	if err != nil {
		if errors.Is(err, usecase.ErrProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "not found",
				Err:     "your item cart product not found",
			}
		}
		if errors.Is(err, usecase.ErrItemCartNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "not found",
				Err:     err.Error(),
			}
		}
		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "updated list cart successfully",
		Data:    cart,
		Err:     nil,
	})
}

func (p *Presenter) DeleteProductCart(c *fiber.Ctx) error {
	id := c.Params("id")
	// _, err = helper.NewUlid(id)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return exception.Err(c, &schema.ErrHttp{
			Code:    fiber.StatusNotFound,
			Message: "NOT FOUND",
			Err:     "item cart not found",
		})
	}

	err = p.CartUsecase.DeleteItemCart(c.Context(), int64(idInt), &usecase.CommonParam{
		UserID: c.Locals("userID").(string),
	})

	if err != nil {
		if errors.Is(err, usecase.ErrProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "BAD REQUEST",
				Err: map[string][]string{
					"product_id": {
						err.Error(),
					},
				},
			}
		}
		if errors.Is(err, usecase.ErrItemCartNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "not found",
				Err:     err.Error(),
			}
		}
		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "deleted list cart successfully",
		Data:    nil,
		Err:     nil,
	})
}

func (p *Presenter) GetProductCart(c *fiber.Ctx) error {
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

	carts, paginate, err := p.CartUsecase.GetItemCartByCart(c.Context(), &usecase.GetCartParam{
		Search: search,
		PaginateParam: usecase.PaginateParam{
			Page:     pageInt,
			PageSize: pageSizeInt,
		},
		CommonParam: usecase.CommonParam{
			UserID: c.Locals("userID").(string),
		},
	})

	if err != nil {
		if errors.Is(err, usecase.ErrProductNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "BAD REQUEST",
				Err: map[string][]string{
					"product_id": {
						err.Error(),
					},
				},
			}
		}
		if errors.Is(err, usecase.ErrItemCartNotFound) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusBadRequest,
				Message: "not found",
				Err:     err.Error(),
			}
		}
		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.Response{
		Code:    200,
		Message: "data carts",
		Data:    carts,
		Err:     nil,
		Paginate: &schema.PaginateRes{
			CurrentPage: paginate.CurrentPage,
			Total:       paginate.Total,
			PageSize:    paginate.PageSize,
			PageTotal:   paginate.PageTotal,
		},
	})
}
