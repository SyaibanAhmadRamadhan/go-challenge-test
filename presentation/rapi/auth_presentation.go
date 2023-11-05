package rapi

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"challenge-test-synapsis/presentation/rapi/exception"
	"challenge-test-synapsis/presentation/rapi/schema"
	"challenge-test-synapsis/usecase"
)

func (p *Presenter) Login(c *fiber.Ctx) error {
	req := new(schema.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return exception.Err(c, err)
	}

	err := req.Validate()
	if err != nil {
		return exception.Err(c, err)
	}

	device := c.Get("User-Agent")
	ip := c.IP()

	auth, err := p.AuthUsecase.Login(c.Context(), &usecase.LoginParam{
		Email:    req.Email,
		Password: req.Password,
		CommonParam: usecase.CommonParam{
			Device: device,
			IP:     ip,
		},
	})

	if err != nil {
		if errors.Is(err, usecase.ErrInvalidEmailOrPass) {
			err = &schema.ErrorHttp{
				Code:    fiber.StatusBadRequest,
				Message: "BAD REQUEST",
				Data:    nil,
				Err:     err.Error(),
			}
		}

		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.ResponseAuth{
		ID:          auth.ID,
		RoleID:      auth.RoleID,
		Username:    auth.Username,
		Email:       auth.Email,
		PhoneNumber: auth.PhoneNumber,
		Token:       auth.Token,
	})
}

func (p *Presenter) Register(c *fiber.Ctx) error {
	req := new(schema.RegisterRequest)

	if err := c.BodyParser(req); err != nil {
		return exception.Err(c, err)
	}

	err := req.Validate()
	if err != nil {
		return exception.Err(c, err)
	}

	device := c.Get("User-Agent")
	ip := c.IP()

	auth, err := p.AuthUsecase.Register(c.Context(), &usecase.RegisterParam{
		RoleID:      1,
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		RememberMe:  false,
		CommonParam: usecase.CommonParam{
			Device: device,
			IP:     ip,
		},
	})

	if err != nil {
		if errors.Is(err, usecase.ErrEmailIsRegistered) {
			err = &schema.ErrorHttp{
				Code:    fiber.StatusBadRequest,
				Message: "BAD REQUEST",
				Data:    nil,
				Err: map[string][]string{
					"email": {
						"email is registered",
					},
				},
			}
		}

		return exception.Err(c, err)
	}

	return c.Status(200).JSON(schema.ResponseAuth{
		ID:          auth.ID,
		RoleID:      auth.RoleID,
		Username:    auth.Username,
		Email:       auth.Email,
		PhoneNumber: auth.PhoneNumber,
		Token:       auth.Token,
	})
}

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
			err = &schema.ErrorHttp{
				Code:    fiber.StatusUnauthorized,
				Message: "UNAUTHORIZATION",
				Data:    nil,
				Err:     err.Error(),
			}
		}

		return exception.Err(c, err)
	}

	c.Set("Authorization", auth.Token)
	c.Locals("role", auth.RoleID)
	return c.Next()
}
