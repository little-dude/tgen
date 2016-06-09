package server

import (
	// "github.com/little-dude/tgen/schemas"
	"strings"
)

// const (
// 	CAPNP_ERROR        = uint8(0)
// 	INTERNAL_ERROR     = uint8(1)
// 	INVALID_PARAMETERS = uint8(2)
// 	UNKNOWN_ERROR      = uint8(3)
// )

type TgenError struct {
	Message string
}

func NewError(msg ...string) error {
	return TgenError{Message: strings.Join(msg, "")}
}

func (err TgenError) Error() string {
	return err.Message
}
