package errors

import "fmt"

type (
	Err struct {
		msg string
	}
)

var (
	NotImplementedError = NewError("not implemented")
)

func NewError(msg string) Err {
	return Err{
		msg: msg,
	}
}

func (e Err) Error() string {
	return e.msg
}

func Wrap(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}
