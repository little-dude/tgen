package fields

import (
	schemas "github.com/little-dude/tgen/capnp"
	"math/rand"
	"zombiezen.com/go/capnproto2"
)

type Field32 struct {
	FirstValue uint32
	FullMask   uint32
	Value      uint32
	Step       uint32
	Count      uint16
	Mode       uint8
	Mask       uint32
}

func (field *Field32) GetValue() []byte {
	value := make([]byte, 4)
	for i := 0; i < 4; i++ {
		value[i] = byte((field.Value >> uint(8*(3-i))) & 255)
	}
	return value
}

func (field *Field32) SetValue(value []byte) {
	value = adjustSliceLength(4, value)
	res := uint32(0)
	for i := 0; i < 4; i++ {
		res += uint32(value[i]) << (8 * uint(3-i))
	}
	field.Value = res & field.FullMask
}

func (field *Field32) GetMask() []byte {
	mask := make([]byte, 4)
	for i := 0; i < 4; i++ {
		mask[i] = byte((field.Mask >> uint(8*(3-i))) & 255)
	}
	return mask
}

func (field *Field32) SetMask(mask []byte) {
	mask = adjustSliceLength(4, mask)
	res := uint32(0)
	for i := 0; i < 4; i++ {
		res += uint32(mask[i]) << (8 * uint(3-i))
	}
	field.Mask = res & field.FullMask
}

func (field *Field32) GetStep() []byte {
	step := make([]byte, 4)
	for i := 0; i < 4; i++ {
		step[i] = byte((field.Step >> uint(8*(3-i))) & 255)
	}
	return step
}

func (field *Field32) SetStep(step []byte) {
	step = adjustSliceLength(4, step)
	res := uint32(0)
	for i := 0; i < 4; i++ {
		res += uint32(step[i]) << (8 * uint(3-i))
	}
	field.Step = res & field.FullMask
}

func (field *Field32) GetCount() uint16 {
	return field.Count
}

func (field *Field32) SetCount(count uint16) {
	if count > 1 {
		field.Count = count
	} else {
		field.Count = 1
	}
}

func (field *Field32) GetMode() uint8 {
	return field.Mode
}

func (field *Field32) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.Mode = mode
	default:
		field.Mode = FIXED
	}
}

func (field *Field32) FromCapnp(capnpField *schemas.Field) {
	value, _ := capnpField.Value()
	field.SetValue(value)

	step, _ := capnpField.Step()
	field.SetStep(step)

	mask, _ := capnpField.Mask()
	field.SetMask(mask)

	field.SetMode(capnpField.Mode())
	field.SetCount(capnpField.Count())
}

func (field *Field32) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)
	capnpField.SetValue(field.GetValue())
	capnpField.SetMode(field.GetMode())
	capnpField.SetStep(field.GetStep())
	capnpField.SetMask(field.GetMask())
	capnpField.SetCount(field.GetCount())
	return capnpField
}

func (field Field32) SetCurrentValue(index uint) {
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

func (field Field32) Increment() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value + field.Step) & field.Mask)
}

func (field Field32) Decrement() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value - field.Step) & field.Mask)
}

func (field Field32) Randomize() {
	field.Value = uint32(rand.Int31()) & field.Mask
}
