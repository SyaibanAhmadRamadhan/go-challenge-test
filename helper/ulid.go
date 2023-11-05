package helper

import (
	"github.com/oklog/ulid/v2"
)

func NewUlid(str string) (strRes string, err error) {
	if str != "" {
		_, err = ulid.Parse(str)
		return
	}

	strRes = ulid.Make().String()
	return
}
