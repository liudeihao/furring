package errors

import (
	"errors"
	"fmt"
)

var (
	As  = errors.As
	Is  = errors.Is
	New = errors.New
)

type InternalError struct {
	err error
	msg string
}

func IsInternalError(err error) bool {
	var internalError InternalError
	ok := errors.As(err, &internalError)
	return ok
}

func (e InternalError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.err.Error())
}

func Internal(err error, msg string) error {
	return InternalError{err: err, msg: msg}
}
