package repository

import (
	"errors"
)

var ErrInvalidColumn = errors.New("column is not exist")
var ErrTxIsNil = errors.New("no Tx Available")
