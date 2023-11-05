package schema

import (
	"fmt"
)

type PaginateRes struct {
	CurrentPage int `json:"current_page"`
	Total       int `json:"total"`
	PageSize    int `json:"page_size"`
	PageTotal   int `json:"page_total"`
}

type Response struct {
	Code     int          `json:"code"`
	Message  string       `json:"message"`
	Data     any          `json:"data"`
	Err      any          `json:"error"`
	Paginate *PaginateRes `json:"paginate,omitempty"`
}

type ErrHttp struct {
	Code    int
	Message string
	Err     any
}

func (e *ErrHttp) Error() string {
	return fmt.Sprintf(e.Message)
}
