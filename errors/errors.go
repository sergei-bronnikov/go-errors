package errors

import (
	"errors"
)

const maxErrorsDepth = 32

//var NotFound = New("not found")
//
//func foo() error {
//	NotFound("file not found")
//}

func New(err string) func(msg ...string) error {
	return func(msg ...string) error {
		message := ""
		if len(msg) > 0 {
			message = msg[0]
		}
		return newErr(err, message)
	}
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, &target)
}

// func Join(errs ...error) error {
// 	return errors.Join(errs...)
// }

func Wrap(err, wrapped error) error {
	return &wrapError{
		err:          err,
		wrappedError: wrapped,
	}
}

func AsString(err error) string {
	var s []byte
	for i := 0; i < maxErrorsDepth; i++ {
		if err == nil {
			break
		}
		if i > 0 {
			s = append(s, []byte("  ")...)
		}
		s = append(s, []byte(err.Error())...)
		switch e := err.(type) {
		case interface {
			Message() string
			Context() string
		}:
			msg := e.Message()
			if msg != "" {
				s = append(s, []byte(" : ")...)
				s = append(s, []byte(msg)...)
			}
			context := e.Context()
			if context != "" {
				s = append(s, []byte(" : ")...)
				s = append(s, []byte(context)...)
			}
		}
		s = append(s, []byte("\n")...)
		if e, ok := err.(interface{ Unwrap() error }); ok {
			err = e.Unwrap()
		} else {
			break
		}
		if i == maxErrorsDepth-1 {
			s = append(s, []byte("  and more...")...)
		}
	}
	return string(s)
}
