package exception

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"

	"challenge-test-synapsis/presentation/rapi/schema"
)

func Err(c *fiber.Ctx, err error) error {
	var (
		errHttp          *schema.ErrorHttp
		errUnmarshalType *json.UnmarshalTypeError
		errSyntak        *json.SyntaxError
	)

	switch {
	case errors.As(err, &errUnmarshalType):
		err = &schema.ErrorHttp{
			Code:    fiber.StatusUnprocessableEntity,
			Message: "UnprocessableEntity",
			Data:    nil,
			Err:     err.Error(),
		}

	case errors.As(err, &errSyntak):
		err = &schema.ErrorHttp{
			Code:    fiber.StatusUnprocessableEntity,
			Message: "unexpected end of json input",
			Data:    nil,
			Err:     err.Error(),
		}
	case errors.Is(err, context.DeadlineExceeded):
		err = &schema.ErrorHttp{
			Code:    fiber.StatusRequestTimeout,
			Message: "request time out",
			Data:    nil,
			Err:     err.Error(),
		}
	}

	ok := errors.As(err, &errHttp)
	if !ok {
		err = &schema.ErrorHttp{
			Code:    fiber.StatusInternalServerError,
			Message: "internal server error",
			Data:    nil,
			Err:     err.Error(),
		}
		errors.As(err, &errHttp)
	}

	return c.Status(err.(*schema.ErrorHttp).Code).JSON(err)

}
