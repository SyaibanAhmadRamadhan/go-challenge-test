package schema

import (
	"fmt"
)

var Required = "field ini tidak boleh dikosongkan"
var MinString = "field ini tidak boleh kurang dari %d"
var MaxString = "field ini tidak boleh lebih dari %d"
var PasswordAndRePassword = "password dan re password tidak sesuai"
var EmailMsg = "email harus menggunakan yourmail@gmail.com"

func MaxMinString(s string, min, max int) string {
	switch {
	case len(s) < min:
		return fmt.Sprintf(MinString, min)
	case len(s) > max:
		return fmt.Sprintf(MaxString, max)
	}

	return ""
}
