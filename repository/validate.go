package repository

import (
	"reflect"
)

func ValidateColumnFromStruct(src any, column string) error {
	val := reflect.ValueOf(src).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("sql")
		if tag == column {
			return nil
		}
	}

	return ErrInvalidColumn
}
