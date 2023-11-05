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
		errHttp          *schema.ErrHttp
		errUnmarshalType *json.UnmarshalTypeError
		errSyntak        *json.SyntaxError
	)

	switch {
	case errors.As(err, &errUnmarshalType):
		err = &schema.ErrHttp{
			Code:    fiber.StatusUnprocessableEntity,
			Message: "UnprocessableEntity",
			Err:     err.Error(),
		}

	case errors.As(err, &errSyntak):
		err = &schema.ErrHttp{
			Code:    fiber.StatusUnprocessableEntity,
			Message: "unexpected end of json input",
			Err:     err.Error(),
		}
	case errors.Is(err, context.DeadlineExceeded):
		err = &schema.ErrHttp{
			Code:    fiber.StatusRequestTimeout,
			Message: "request time out",
			Err:     err.Error(),
		}
	}

	ok := errors.As(err, &errHttp)
	if !ok {
		err = &schema.ErrHttp{
			Code:    fiber.StatusInternalServerError,
			Message: "internal server error",
			Err:     err.Error(),
		}
		errors.As(err, &errHttp)
	}

	response := schema.Response{
		Code:    err.(*schema.ErrHttp).Code,
		Message: err.(*schema.ErrHttp).Message,
		Data:    nil,
		Err:     err.(*schema.ErrHttp).Err,
	}

	return c.Status(response.Code).JSON(response)

}
