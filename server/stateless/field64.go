package stateless

import (
	"math/rand"
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
	value := make([]byte, 8)
	for i := 0; i < 8; i++ {
		value[i] = byte((field.Value >> uint(8*(7-i))) & 255)
	}
	return value
}

func (field *Field64) SetValue(value []byte) {
	value = adjustSliceLength(2, value)
	res := uint64(0)
	for i := 0; i < 8; i++ {
		res += uint64(value[i]) << (8 * uint(7-i))
	}
	field.Value = res & field.FullMask
}

func (field *Field64) GetMask() []byte {
	mask := make([]byte, 8)
	for i := 0; i < 8; i++ {
		mask[i] = byte((field.Mask >> uint(8*(7-i))) & 255)
	}
	return mask
}

func (field *Field64) SetMask(mask []byte) {
	mask = adjustSliceLength(2, mask)
	res := uint64(0)
	for i := 0; i < 8; i++ {
		res += uint64(mask[i]) << (8 * uint(7-i))
	}
	field.Mask = res & field.FullMask
}

func (field *Field64) GetStep() []byte {
	step := make([]byte, 8)
	for i := 0; i < 8; i++ {
		step[i] = byte((field.Step >> uint(8*(7-i))) & 255)
	}
	return step
}

func (field *Field64) SetStep(step []byte) {
	step = adjustSliceLength(8, step)
	res := uint64(0)
	for i := 0; i < 8; i++ {
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

func (field *Field64) SetCurrentValue(index uint) {
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

func (field *Field64) Increment() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value + field.Step) & field.Mask)
}

func (field *Field64) Decrement() {
	field.Value = (field.Value & (^field.Mask)) | ((field.Value - field.Step) & field.Mask)
}

func (field *Field64) Randomize() {
	field.Value = uint64(rand.Int63()) & field.Mask
}
