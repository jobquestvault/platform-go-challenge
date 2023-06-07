package sys

import "fmt"

type (
	Err struct {
		msg string
	}
)

func NewError(msg string) Err {
	return Err{
		msg: msg,
	}
}

func (e Err) String() string {
	return e.msg
}

func Wrap(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}
