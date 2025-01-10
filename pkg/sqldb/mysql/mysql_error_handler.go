package xmysql

import (
	"errors"
	"github.com/go-sql-driver/mysql"
)

const (
	DuplicatePrimaryKeyErrorCode ErrorHandlerCode = 1062
	UniqueViolationErrorCode     ErrorHandlerCode = 1169
)

type ErrorHandlerCode uint16
type ErrorHandlers map[ErrorHandlerCode]ErrorHandlerFunc
type ErrorHandlerFunc func(resource interface{}, err error) error

type ErrorHandler struct {
	handlers ErrorHandlers
}

func NewErrorHandler(handlers map[ErrorHandlerCode]ErrorHandlerFunc) *ErrorHandler {
	return &ErrorHandler{handlers: handlers}
}

func (eh *ErrorHandler) Handle(resource interface{}, err error) error {
	var mysqlErr *mysql.MySQLError

	if errors.As(err, &mysqlErr) {
		errCode := ErrorHandlerCode(mysqlErr.Number)
		if handler, exists := eh.handlers[errCode]; exists {
			return handler(resource, err)
		}
	}

	return err
}
