package stateless

import (
	"math/rand"
)

type LongField struct {
	FirstValue []byte
	FullMask   []byte
	Value      []byte
	Step       []byte
	Count      uint16
	Mode       uint8
	Mask       []byte
}

func (field *LongField) GetValue() []byte {
	value := make([]byte, len(field.Value))
	copy(value, field.Value)
	return value
}

func (field *LongField) SetValue(value []byte) {
	value = adjustSliceLength(len(field.FullMask), value)
	for i := 0; i < len(value); i++ {
		value[i] = value[i] & field.FullMask[i]
	}
	field.Value = value
}

func (field *LongField) GetMask() []byte {
	mask := make([]byte, len(field.Mask))
	copy(mask, field.Mask)
	return mask
}

func (field *LongField) SetMask(mask []byte) {
	mask = adjustSliceLength(len(field.FullMask), mask)
	for i := 0; i < len(mask); i++ {
		mask[i] = mask[i] & field.FullMask[i]
	}
	field.Mask = mask
}

func (field *LongField) GetStep() []byte {
	step := make([]byte, len(field.Step))
	copy(step, field.Step)
	return step
}

func (field *LongField) SetStep(step []byte) {
	step = adjustSliceLength(len(field.FullMask), step)
	for i := 0; i < len(step); i++ {
		step[i] = step[i] & field.FullMask[i]
	}
	field.Step = step
}

func (field *LongField) GetCount() uint16 {
	return field.Count
}

func (field *LongField) SetCount(count uint16) {
	if count > 1 {
		field.Count = count
	} else {
		field.Count = 1
	}
}

func (field *LongField) GetMode() uint8 {
	return field.Mode
}

func (field *LongField) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.Mode = mode
	default:
		field.Mode = FIXED
	}
}

func (field *LongField) SetCurrentValue(index uint) {
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
	default:
		return
	}
}

func (field *LongField) Randomize() {
	for i := len(field.Value) - 1; i >= 0; i-- {
		field.Value[i] = byte(rand.Intn(256))
	}
}

func (field *LongField) Decrement() {
	overflow := false
	var newValue byte
	for i := len(field.Value) - 1; i >= 0; i-- {
		if field.Mask[i] == 0 {
			overflow = false
			continue
		}
		if overflow == true {
			newValue, overflow = sub(field.Value[i], 1)
		} else {
			newValue = field.Value[i]
		}
		if overflow == true {
			newValue, _ = sub(newValue, field.Step[i])
		} else {
			newValue, overflow = sub(newValue, field.Step[i])
		}
		field.Value[i] = newValue & field.Mask[i]
	}
}

func (field *LongField) Increment() {
	overflow := byte(0)
	var newValue, tmpOverflow byte
	for i := len(field.Value) - 1; i >= 0; i-- {
		if field.Mask[i] == 0 {
			overflow = 0
			continue
		}
		if overflow > 0 {
			newValue, tmpOverflow = add(field.Value[i], overflow)
		} else {
			newValue = field.Value[i]
			tmpOverflow = 0
		}
		newValue, overflow = add(newValue, field.Step[i])
		field.Value[i] = newValue & field.Mask[i]
		overflow += tmpOverflow
	}
}

// Modulo:
// a = q*b + r (where r=a%b)  => a%b = a - q*b
// in practice a/b = q
// so a%b = a - (a/b)*b
//
// now how to perform division on []byte?
//
// 	9876 | 5
// 	----------
//  9	 |
// -5 	 | 1
// =4	 |
//  48	 |
// -45 	 |  9
// = 3	 |
//   37  |
//  -35	 |   7
//  = 2	 |
//    26 |
//   -25 |    5
//   = 1
//
// ===> 9876 = 1975*5 + 1
//
// This shows how to perform division on an array:
// [9, 8, 7, 6] / [5] = [9/5, (8+9%5)/5, (7+(8+9%5)%5)/5, (6+(7+(8+9%5)%5)%5)/5] + 9876 % 5
//
// this is easy to implement but it does not work for divisors with more than
// one digit, and I could not figure out how to generalize
//
// https://en.wikipedia.org/wiki/Long_division
// http://courses.cs.vt.edu/~cs1104/BuildingBlocks/divide.030.html
// http://stackoverflow.com/questions/3199727/how-to-implement-long-division-for-enormous-numbers-bignums
