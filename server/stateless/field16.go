package stateless

import (
	"math/rand"
)

type Field16 struct {
	FirstValue uint16
	FullMask   uint16
	Value      uint16
	Step       uint16
	Count      uint16
	Mode       uint8
	Mask       uint16
}

func (field *Field16) GetValue() []byte {
	return []byte{byte(field.Value >> 8), byte(field.Value & 255)}
}

func (field *Field16) SetValue(value []byte) {
	value = adjustSliceLength(2, value)
	res := uint16(0)
	for i := 0; i < 2; i++ {
		res += uint16(value[i]) << (8 * uint(1-i))
	}
	field.Value = res & field.FullMask
}

func (field *Field16) GetMask() []byte {
	return []byte{byte(field.Mask >> 8), byte(field.Mask & 255)}
}

func (field *Field16) SetMask(mask []byte) {
	mask = adjustSliceLength(2, mask)
	res := uint16(0)
	for i := 0; i < 2; i++ {
		res += uint16(mask[i]) << (8 * uint(1-i))
	}
	field.Mask = res & field.FullMask
}

func (field *Field16) GetStep() []byte {
	return []byte{byte(field.Step >> 8), byte(field.Step & 255)}
}

func (field *Field16) SetStep(step []byte) {
	step = adjustSliceLength(2, step)
	res := uint16(0)
	for i := 0; i < 2; i++ {
		res += uint16(step[i]) << (8 * uint(1-i))
	}
	field.Step = res & field.FullMask
}

func (field *Field16) GetCount() uint16 {
	return field.Count
}

func (field *Field16) SetCount(count uint16) {
	if count > field.FullMask {
		field.Count = field.FullMask + 1
	} else if count == 0 {
		field.Count = 1
	} else {
		field.Count = count
	}
}

func (field *Field16) GetMode() uint8 {
	return field.Mode
}

func (field *Field16) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.Mode = mode
	default:
		field.Mode = FIXED
	}
}

func (field *Field16) SetCurrentValue(index uint) {
	if index%uint(field.Count) == 0 && field.Mode != RANDOMIZE {
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

func (field *Field16) Increment() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value + field.Step) & field.Mask)
}

func (field *Field16) Decrement() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value - field.Step) & field.Mask)
}

func (field *Field16) Randomize() {
	field.Value = uint16(rand.Int31()) & field.Mask
}
