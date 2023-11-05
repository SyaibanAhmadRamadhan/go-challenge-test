package schema

import (
	"fmt"
)

type ErrorHttp struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Err     any    `json:"error,omitempty"`
}

func (e *ErrorHttp) Error() string {
	return fmt.Sprintf(e.Message)
}
