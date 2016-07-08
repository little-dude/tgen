package errors

import "strings"

type TgenError struct {
	Message string
}

func New(msg ...string) error {
	return TgenError{Message: strings.Join(msg, " ")}
}

func (err TgenError) Error() string {
	return err.Message
}
