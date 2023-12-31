package schema

import (
	"fmt"
)

var Required = "This field is required"
var MinString = "This field cannot be less than %d"
var MaxString = "This field cannot be more than %d"
var MinInt = "cannot be less than %d"
var MaxInt = "cannot be more than %d"
var PasswordAndRePassword = "password and re-password do not match"
var EmailMsg = "email must be yourmail@gmail.com"

func MaxMinString(s string, min, max int) string {
	switch {
	case len(s) < min:
		return fmt.Sprintf(MinString, min)
	case len(s) > max:
		return fmt.Sprintf(MaxString, max)
	}

	return ""
}

func MaxMinInt(i, min, max int) string {
	switch {
	case i < min:
		return fmt.Sprintf(MinInt, min)
	case i > max:
		return fmt.Sprintf(MaxInt, max)
	}

	return ""
}
