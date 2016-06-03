package fields

import (
	schemas "github.com/little-dude/tgen/capnp"
	"math/rand"
	"zombiezen.com/go/capnproto2"
)

type Field8 struct {
	FirstValue uint8
	FullMask   uint8
	Value      uint8
	Step       uint8
	Count      uint8
	Mode       uint8
	Mask       uint8
}

func (field *Field8) GetValue() []byte {
	return []byte{byte(field.Value)}
}

func (field *Field8) SetValue(value []byte) {
	field.Value = adjustSliceLength(1, value)[0] & field.FullMask
}

func (field *Field8) GetMask() []byte {
	return []byte{byte(field.Mask)}
}

func (field *Field8) SetMask(mask []byte) {
	field.Mask = adjustSliceLength(1, mask)[0] & field.FullMask
}

func (field *Field8) GetStep() []byte {
	return []byte{byte(field.Step)}
}

func (field *Field8) SetStep(step []byte) {
	field.Step = adjustSliceLength(1, step)[0] & field.FullMask
}

func (field *Field8) GetCount() uint64 {
	return uint64(field.Count)
}

func (field *Field8) SetCount(count uint64) {
	if count > 1 {
		field.Count = uint8(count % uint64(field.FullMask))
	} else {
		field.Count = uint8(1)
	}
}

func (field *Field8) GetMode() uint8 {
	return field.Mode
}

func (field *Field8) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.Mode = mode
	default:
		field.Mode = FIXED
	}
}

func (field *Field8) FromCapnp(capnpField *schemas.Field) {
	value, _ := capnpField.Value()
	field.SetValue(value)

	step, _ := capnpField.Step()
	field.SetStep(step)

	mask, _ := capnpField.Mask()
	field.SetMask(mask)

	field.SetMode(capnpField.Mode())
	field.SetCount(capnpField.Count())
}

func (field *Field8) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)
	capnpField.SetValue(field.GetValue())
	capnpField.SetMode(field.GetMode())
	capnpField.SetStep(field.GetStep())
	capnpField.SetMask(field.GetMask())
	capnpField.SetCount(field.GetCount())
	return capnpField
}

func (field *Field8) Increment() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value + field.Step) & field.Mask)
}

func (field *Field8) Decrement() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value - field.Step) & field.Mask)
}

func (field *Field8) Randomize() {
	field.Value = uint8(rand.Int31()) & field.Mask & field.FullMask
}

func (field *Field8) SetCurrentValue(index uint) {
	if index%uint(field.Count) == 0 {
		field.Value = field.FirstValue
		return
	}
	switch field.Mode {
	case INCREMENT:
		field.Increment()
	case DECREMENT:
		field.Decrement()
	case RANDOMIZE:
		field.Randomize()
	}
}
