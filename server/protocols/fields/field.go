package fields

import (
	"fmt"
	schemas "github.com/little-dude/tgen/capnp"
	// "math/rand"
	"zombiezen.com/go/capnproto2"
)

const (
	FIXED     = iota
	AUTO      = iota
	INCREMENT = iota
	DECREMENT = iota
	RANDOMIZE = iota
)

type Field interface {
	SetValue([]byte)
	SetStep([]byte)
	SetMask([]byte)
	SetCount(uint16)
	SetMode(uint8)
	// SetFullMask([]byte)

	GetValue() []byte
	GetStep() []byte
	GetMask() []byte
	GetCount() uint16
	GetMode() uint8
	// GetFullMask() []byte

	Increment()
	Decrement()
	Randomize()

	SetCurrentValue(uint)
	ToCapnp(*capnp.Segment) schemas.Field
	FromCapnp(*schemas.Field)
}

func toString(field Field) string {
	return fmt.Sprint(
		"{value: ", field.GetValue(),
		", mask: ", field.GetMask(),
		", step: ", field.GetStep(),
		", mode: ", field.GetMode(),
		", count: ", field.GetCount(),
		"}")
}

func (field Field8) String() string {
	return toString(&field)
}
func (field Field16) String() string {
	return toString(&field)
}
func (field Field32) String() string {
	return toString(&field)
}
func (field Field64) String() string {
	return toString(&field)
}
func (field LongField) String() string {
	return toString(&field)
}
