package stateless

import (
	"fmt"
	"github.com/little-dude/tgen/server/schemas"
	// "math/rand"
	// "zombiezen.com/go/capnproto2"
)

const (
	FIXED     = iota
	AUTO      = iota
	INCREMENT = iota
	DECREMENT = iota
	RANDOMIZE = iota
)

type _Field interface {
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
}

type Field interface {
	ToCapnp(*schemas.Field) error
	FromCapnp(*schemas.Field) error
}

func CerealizeField(f _Field, capnpField *schemas.Field) error {
	e := capnpField.SetValue(f.GetValue())
	if e != nil {
		return e
	}

	e = capnpField.SetStep(f.GetStep())
	if e != nil {
		return e
	}

	e = capnpField.SetMask(f.GetMask())
	if e != nil {
		return e
	}

	capnpField.SetMode(f.GetMode())
	capnpField.SetCount(f.GetCount())

	return nil
}

func DecerealizeField(f _Field, capnpField *schemas.Field) error {
	value, e := capnpField.Value()
	if e != nil {
		return e
	}

	mask, e := capnpField.Mask()
	if e != nil {
		return e
	}

	step, e := capnpField.Step()
	if e != nil {
		return e
	}

	f.SetValue(value)
	f.SetStep(step)
	f.SetMask(mask)
	f.SetMode(capnpField.Mode())
	f.SetCount(capnpField.Count())
	return nil
}

func (field *Field8) FromCapnp(capnpField *schemas.Field) error {
	return DecerealizeField(field, capnpField)
}

func (field *Field8) ToCapnp(capnpField *schemas.Field) error {
	return CerealizeField(field, capnpField)
}

func (field *Field16) FromCapnp(capnpField *schemas.Field) error {
	return DecerealizeField(field, capnpField)
}

func (field *Field16) ToCapnp(capnpField *schemas.Field) error {
	return CerealizeField(field, capnpField)
}

func (field *Field32) FromCapnp(capnpField *schemas.Field) error {
	return DecerealizeField(field, capnpField)
}

func (field *Field32) ToCapnp(capnpField *schemas.Field) error {
	return CerealizeField(field, capnpField)
}

func (field *Field64) FromCapnp(capnpField *schemas.Field) error {
	return DecerealizeField(field, capnpField)
}

func (field *Field64) ToCapnp(capnpField *schemas.Field) error {
	return CerealizeField(field, capnpField)
}

func (field *LongField) FromCapnp(capnpField *schemas.Field) error {
	return DecerealizeField(field, capnpField)
}

func (field *LongField) ToCapnp(capnpField *schemas.Field) error {
	return CerealizeField(field, capnpField)
}

func toString(field Field) string {
	f := field.(_Field)
	return fmt.Sprint(
		"{value: ", f.GetValue(),
		", mask: ", f.GetMask(),
		", step: ", f.GetStep(),
		", mode: ", f.GetMode(),
		", count: ", f.GetCount(),
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
