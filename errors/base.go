package errors

import (
	"errors"
	"fmt"
	"runtime"
)

type BaseError struct {
	err     string
	msg     string
	context string
}

func (e *BaseError) Error() string {
	return e.err
}

func (e *BaseError) Is(target error) bool {
	return e.Error() == target.Error()
}

func (e *BaseError) Message() string {
	return e.msg
}

func (e *BaseError) Context() string {
	return e.context
}

func AsBaseError(err error) *BaseError {
	var baseError *BaseError
	ok := errors.As(err, &baseError)
	if ok {
		return baseError
	}
	return nil
}

func newErr(err, msg string) error {
	return &BaseError{
		err:     err,
		msg:     msg,
		context: getContext(),
	}
}

func getContext() string {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		funcN := runtime.FuncForPC(pc).Name()
		return fmt.Sprintf("%s[%d] - %s", file, line, funcN)
	}
	return "unknown"
}
