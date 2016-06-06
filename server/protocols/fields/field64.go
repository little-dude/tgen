package fields

import (
	schemas "github.com/little-dude/tgen/capnp"
	"math/rand"
	"zombiezen.com/go/capnproto2"
)

type Field64 struct {
	FirstValue uint64
	FullMask   uint64
	Value      uint64
	Step       uint64
	Count      uint16
	Mode       uint8
	Mask       uint64
}

func (field *Field64) GetValue() []byte {
	return []byte{byte(field.Value >> 8), byte(field.Value & 255)}
}

func (field *Field64) SetValue(value []byte) {
	value = adjustSliceLength(2, value)
	res := uint64(0)
	for i := 0; i < 2; i++ {
		res += uint64(value[i]) << (8 * uint(7-i))
	}
	field.Value = res & field.FullMask
}

func (field *Field64) GetMask() []byte {
	return []byte{byte(field.Mask >> 8), byte(field.Mask & 255)}
}

func (field *Field64) SetMask(mask []byte) {
	mask = adjustSliceLength(2, mask)
	res := uint64(0)
	for i := 0; i < 2; i++ {
		res += uint64(mask[i]) << (8 * uint(7-i))
	}
	field.Mask = res & field.FullMask
}

func (field *Field64) GetStep() []byte {
	return []byte{byte(field.Step >> 8), byte(field.Step & 255)}
}

func (field *Field64) SetStep(step []byte) {
	step = adjustSliceLength(2, step)
	res := uint64(0)
	for i := 0; i < 2; i++ {
		res += uint64(step[i]) << (8 * uint(7-i))
	}
	field.Step = res & field.FullMask
}

func (field *Field64) GetCount() uint16 {
	return field.Count
}

func (field *Field64) SetCount(count uint16) {
	if count > 1 {
		field.Count = count
	} else {
		field.Count = 1
	}
}

func (field *Field64) GetMode() uint8 {
	return field.Mode
}

func (field *Field64) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.Mode = mode
	default:
		field.Mode = FIXED
	}
}

func (field *Field64) FromCapnp(capnpField *schemas.Field) {
	value, _ := capnpField.Value()
	field.SetValue(value)

	step, _ := capnpField.Step()
	field.SetStep(step)

	mask, _ := capnpField.Mask()
	field.SetMask(mask)

	field.SetMode(capnpField.Mode())
	field.SetCount(capnpField.Count())
}

func (field *Field64) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)
	capnpField.SetValue(field.GetValue())
	capnpField.SetMode(field.GetMode())
	capnpField.SetStep(field.GetStep())
	capnpField.SetMask(field.GetMask())
	capnpField.SetCount(field.GetCount())
	return capnpField
}
func (field Field64) SetCurrentValue(index uint) {
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

func (field Field64) Increment() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value + field.Step) & field.Mask)
}

func (field Field64) Decrement() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value - field.Step) & field.Mask)
}

func (field Field64) Randomize() {
	field.Value = uint64(rand.Int63()) & field.Mask
}
