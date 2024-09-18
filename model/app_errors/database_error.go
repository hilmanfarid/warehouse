package app_errors

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

var (
	mysqlErr *mysql.MySQLError
)

func IsMySQLDuplicateKey(err error) bool {
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return true
	}
	return false
}

func IsMySQLReferenceError(err error) bool {
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		return true
	}
	return false
}

func NewMySQLDuplicateKey(err error) error {
	errDetails := NewErrorDetails("record already exists", err.Error(), "422")
	return NewBaseError(errDetails)
}

func NewMySQLNotFound(err error) error {
	if err == nil {
		err = errors.New("not found")
	}
	errDetails := NewErrorDetails("not found", err.Error(), "404")
	return NewBaseError(errDetails)
}

func NewMySQLReferenceError(err error) error {
	errDetails := NewErrorDetails("", err.Error(), "404")
	return NewBaseError(errDetails)
}
