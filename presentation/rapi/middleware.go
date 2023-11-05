package rapi

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"challenge-test-synapsis/presentation/rapi/exception"
	"challenge-test-synapsis/presentation/rapi/schema"
	"challenge-test-synapsis/usecase"
)

func (p *Presenter) Otorisasi(c *fiber.Ctx) error {
	device := c.Get("User-Agent")
	token := c.Get("Authorization")
	userID := c.Get("User-Id")
	ip := c.IP()

	common := usecase.CommonParam{
		Device: device,
		IP:     ip,
		UserID: userID,
	}
	auth, err := p.AuthUsecase.Otorisasi(c.Context(), token, &common)

	if err != nil {
		if errors.Is(err, usecase.ErrInvalidToken) {
			err = &schema.ErrHttp{
				Code:    fiber.StatusUnauthorized,
				Message: "UNAUTHORIZATION",
				Err:     err.Error(),
			}
		}

		return exception.Err(c, err)
	}

	c.Set("Authorization", auth.Token)
	c.Locals("role", auth.RoleID)
	c.Locals("userID", userID)
	return c.Next()
}

func MustBeAdmin(c *fiber.Ctx) error {
	roleID := c.Locals("role").(int)
	if roleID != 2 {
		err := &schema.ErrHttp{
			Code:    fiber.StatusForbidden,
			Message: "FORBIDDEN",
			Err:     "you don't have access",
		}
		return exception.Err(c, err)
	}

	return c.Next()
}

func CheckLogin(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	userID := c.Get("User-Id")

	if token != "" || userID != "" {
		err := &schema.ErrHttp{
			Code:    fiber.StatusForbidden,
			Message: "FORBIDDEN",
			Err:     "you are logged in",
		}

		return exception.Err(c, err)
	}

	return c.Next()
}
