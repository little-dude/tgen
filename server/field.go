package server

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

import (
	"fmt"
	schemas "github.com/little-dude/tgen/capnp"
	"math/rand"
	"zombiezen.com/go/capnproto2"
)

const (
	FIXED     = iota
	AUTO      = iota
	INCREMENT = iota
	DECREMENT = iota
	RANDOMIZE = iota
)

func adjustSliceLength(length int, slice []byte) []byte {
	l := len(slice)
	if l == length {
		return slice
	}
	if l < length {
		// left padding with 0s
		return append(make([]byte, length-l), slice...)
	}
	// l > length so we truncate it, taking the right-most words
	return slice[l-length:]
}

type Field interface {
	SetValue([]byte)
	SetStep([]byte)
	SetMask([]byte)
	SetCount(uint64)
	SetMode(uint8)

	GetValue() []byte
	GetStep() []byte
	GetMask() []byte
	GetCount() uint64
	GetMode() uint8

	Increment()
	Decrement()
	Randomize()

	SetCurrentValue(uint)
	ToCapnp(*capnp.Segment) schemas.Field
	FromCapnp(*schemas.Field)
}

type Field8 struct {
	firstValue uint8
	fullMask   uint8
	value      uint8
	step       uint8
	count      uint8
	mode       uint8
	mask       uint8
}

func (field *Field8) GetValue() []byte {
	return []byte{byte(field.value)}
}

func (field *Field8) SetValue(value []byte) {
	field.value = adjustSliceLength(1, value)[0] & field.fullMask
}

func (field *Field8) GetMask() []byte {
	return []byte{byte(field.mask)}
}

func (field *Field8) SetMask(mask []byte) {
	field.mask = adjustSliceLength(1, mask)[0] & field.fullMask
}

func (field *Field8) GetStep() []byte {
	return []byte{byte(field.step)}
}

func (field *Field8) SetStep(step []byte) {
	field.step = adjustSliceLength(1, step)[0] & field.fullMask
}

func (field *Field8) GetCount() uint64 {
	return uint64(field.count)
}

func (field *Field8) SetCount(count uint64) {
	if count > 1 {
		field.count = uint8(count % uint64(field.fullMask))
	} else {
		field.count = uint8(1)
	}
}

func (field *Field8) GetMode() uint8 {
	return field.mode
}

func (field *Field8) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.mode = mode
	default:
		Error.Println("Invalid mode: ", mode, " setting mode to ", FIXED, " (fixed)")
		field.mode = FIXED
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
	field.value = (field.value & (^field.mask)) | ((field.value + field.step) & field.mask)
}

func (field *Field8) Decrement() {
	field.value = (field.value & (^field.mask)) | ((field.value - field.step) & field.mask)
}

func (field *Field8) Randomize() {
	field.value = uint8(rand.Int31()) & field.mask & field.fullMask
}

func (field *Field8) SetCurrentValue(index uint) {
	if index%uint(field.count) == 0 {
		field.value = field.firstValue
		return
	}
	switch field.mode {
	case INCREMENT:
		field.Increment()
	case DECREMENT:
		field.Decrement()
	case RANDOMIZE:
		field.Randomize()
	}
}

type Field16 struct {
	firstValue uint16
	fullMask   uint16
	value      uint16
	step       uint16
	count      uint16
	mode       uint8
	mask       uint16
}

func (field *Field16) GetValue() []byte {
	return []byte{byte(field.value >> 8), byte(field.value & 255)}
}

func (field *Field16) SetValue(value []byte) {
	value = adjustSliceLength(2, value)
	res := uint16(0)
	for i := 0; i < 2; i++ {
		res += uint16(value[i]) << (8 * uint(1-i))
	}
	field.value = res & field.fullMask
}

func (field *Field16) GetMask() []byte {
	return []byte{byte(field.mask >> 8), byte(field.mask & 255)}
}

func (field *Field16) SetMask(mask []byte) {
	mask = adjustSliceLength(2, mask)
	res := uint16(0)
	for i := 0; i < 2; i++ {
		res += uint16(mask[i]) << (8 * uint(1-i))
	}
	field.mask = res & field.fullMask
}

func (field *Field16) GetStep() []byte {
	return []byte{byte(field.step >> 8), byte(field.step & 255)}
}

func (field *Field16) SetStep(step []byte) {
	step = adjustSliceLength(2, step)
	res := uint16(0)
	for i := 0; i < 2; i++ {
		res += uint16(step[i]) << (8 * uint(1-i))
	}
	field.step = res & field.fullMask
}

func (field *Field16) GetCount() uint64 {
	return uint64(field.count)
}

func (field *Field16) SetCount(count uint64) {
	if count > 1 {
		field.count = uint16(count % uint64(field.fullMask))
	} else {
		field.count = uint16(1)
	}
}

func (field *Field16) GetMode() uint8 {
	return field.mode
}

func (field *Field16) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.mode = mode
	default:
		Error.Println("Invalid mode: ", mode, " setting mode to ", FIXED, " (fixed)")
		field.mode = FIXED
	}
}

func (field *Field16) FromCapnp(capnpField *schemas.Field) {
	value, _ := capnpField.Value()
	field.SetValue(value)

	step, _ := capnpField.Step()
	field.SetStep(step)

	mask, _ := capnpField.Mask()
	field.SetMask(mask)

	field.SetMode(capnpField.Mode())
	field.SetCount(capnpField.Count())
}

func (field *Field16) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)
	capnpField.SetValue(field.GetValue())
	capnpField.SetMode(field.GetMode())
	capnpField.SetStep(field.GetStep())
	capnpField.SetMask(field.GetMask())
	capnpField.SetCount(field.GetCount())
	return capnpField
}

func (field Field16) SetCurrentValue(index uint) {
	if index%uint(field.count) == 0 {
		field.value = field.firstValue
		return
	}
	switch field.mode {
	case INCREMENT:
		field.Increment()
	case DECREMENT:
		field.Decrement()
	case RANDOMIZE:
		field.Randomize()
	}
}

func (field Field16) Increment() {
	field.value = (field.value & (^field.mask)) | ((field.value + field.step) & field.mask)
}

func (field Field16) Decrement() {
	field.value = (field.value & (^field.mask)) | ((field.value - field.step) & field.mask)
}

func (field Field16) Randomize() {
	field.value = uint16(rand.Int31()) & field.mask
}

type Field32 struct {
	firstValue uint32
	fullMask   uint32
	value      uint32
	step       uint32
	count      uint32
	mode       uint8
	mask       uint32
}

func (field *Field32) GetValue() []byte {
	value := make([]byte, 4)
	for i := 0; i < 4; i++ {
		value[i] = byte((field.value >> uint(8*(3-i))) & 255)
	}
	return value
}

func (field *Field32) SetValue(value []byte) {
	value = adjustSliceLength(4, value)
	res := uint32(0)
	for i := 0; i < 4; i++ {
		res += uint32(value[i]) << (8 * uint(3-i))
	}
	field.value = res & field.fullMask
}

func (field *Field32) GetMask() []byte {
	mask := make([]byte, 4)
	for i := 0; i < 4; i++ {
		mask[i] = byte((field.mask >> uint(8*(3-i))) & 255)
	}
	return mask
}

func (field *Field32) SetMask(mask []byte) {
	mask = adjustSliceLength(4, mask)
	res := uint32(0)
	for i := 0; i < 4; i++ {
		res += uint32(mask[i]) << (8 * uint(3-i))
	}
	field.mask = res & field.fullMask
}

func (field *Field32) GetStep() []byte {
	step := make([]byte, 4)
	for i := 0; i < 4; i++ {
		step[i] = byte((field.step >> uint(8*(3-i))) & 255)
	}
	return step
}

func (field *Field32) SetStep(step []byte) {
	step = adjustSliceLength(4, step)
	res := uint32(0)
	for i := 0; i < 4; i++ {
		res += uint32(step[i]) << (8 * uint(3-i))
	}
	field.step = res & field.fullMask
}

func (field *Field32) GetCount() uint64 {
	return uint64(field.count)
}

func (field *Field32) SetCount(count uint64) {
	if count > 1 {
		field.count = uint32(count % uint64(field.fullMask))
	} else {
		field.count = uint32(1)
	}
}

func (field *Field32) GetMode() uint8 {
	return field.mode
}

func (field *Field32) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.mode = mode
	default:
		Error.Println("Invalid mode: ", mode, " setting mode to ", FIXED, " (fixed)")
		field.mode = FIXED
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
	if uint32(index)%field.count == 0 {
		field.value = field.firstValue
		return
	}
	switch field.mode {
	case INCREMENT:
		field.Increment()
	case DECREMENT:
		field.Decrement()
	case RANDOMIZE:
		field.Randomize()
	}
}

func (field Field32) Increment() {
	field.value = (field.value & (^field.mask)) | ((field.value + field.step) & field.mask)
}

func (field Field32) Decrement() {
	field.value = (field.value & (^field.mask)) | ((field.value - field.step) & field.mask)
}

func (field Field32) Randomize() {
	field.value = uint32(rand.Int31()) & field.mask
}

type Field64 struct {
	firstValue uint64
	fullMask   uint64
	value      uint64
	step       uint64
	count      uint64
	mode       uint8
	mask       uint64
}

func (field *Field64) GetValue() []byte {
	return []byte{byte(field.value >> 8), byte(field.value & 255)}
}

func (field *Field64) SetValue(value []byte) {
	value = adjustSliceLength(2, value)
	res := uint64(0)
	for i := 0; i < 2; i++ {
		res += uint64(value[i]) << (8 * uint(7-i))
	}
	field.value = res & field.fullMask
}

func (field *Field64) GetMask() []byte {
	return []byte{byte(field.mask >> 8), byte(field.mask & 255)}
}

func (field *Field64) SetMask(mask []byte) {
	mask = adjustSliceLength(2, mask)
	res := uint64(0)
	for i := 0; i < 2; i++ {
		res += uint64(mask[i]) << (8 * uint(7-i))
	}
	field.mask = res & field.fullMask
}

func (field *Field64) GetStep() []byte {
	return []byte{byte(field.step >> 8), byte(field.step & 255)}
}

func (field *Field64) SetStep(step []byte) {
	step = adjustSliceLength(2, step)
	res := uint64(0)
	for i := 0; i < 2; i++ {
		res += uint64(step[i]) << (8 * uint(7-i))
	}
	field.step = res & field.fullMask
}

func (field *Field64) GetCount() uint64 {
	return uint64(field.count)
}

func (field *Field64) SetCount(count uint64) {
	if count > 1 {
		field.count = uint64(count % uint64(field.fullMask))
	} else {
		field.count = uint64(1)
	}
}

func (field *Field64) GetMode() uint8 {
	return field.mode
}

func (field *Field64) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.mode = mode
	default:
		Error.Println("Invalid mode: ", mode, " setting mode to ", FIXED, " (fixed)")
		field.mode = FIXED
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
	if uint64(index)%field.count == 0 {
		field.value = field.firstValue
		return
	}
	switch field.mode {
	case INCREMENT:
		field.Increment()
	case DECREMENT:
		field.Decrement()
	case RANDOMIZE:
		field.Randomize()
	}
}

func (field Field64) Increment() {
	field.value = (field.value & (^field.mask)) | ((field.value + field.step) & field.mask)
}

func (field Field64) Decrement() {
	field.value = (field.value & (^field.mask)) | ((field.value - field.step) & field.mask)
}

func (field Field64) Randomize() {
	field.value = uint64(rand.Int63()) & field.mask
}

type LongField struct {
	firstValue []byte
	fullMask   []byte
	value      []byte
	step       []byte
	count      uint64 // does not really make sense to generate more than 2^64 different packets
	mode       uint8
	mask       []byte
}

func (field *LongField) GetValue() []byte {
	return field.value
}

func (field *LongField) SetValue(value []byte) {
	value = adjustSliceLength(len(field.fullMask), value)
	for i := 0; i < len(value); i++ {
		value[i] = value[i] & field.fullMask[i]
	}
	field.value = value
}

func (field *LongField) GetMask() []byte {
	return field.mask
}

func (field *LongField) SetMask(mask []byte) {
	mask = adjustSliceLength(len(field.fullMask), mask)
	for i := 0; i < len(mask); i++ {
		mask[i] = mask[i] & field.fullMask[i]
	}
	field.mask = mask
}

func (field *LongField) GetStep() []byte {
	return field.step
}

func (field *LongField) SetStep(step []byte) {
	Error.Println("input", step)
	step = adjustSliceLength(len(field.fullMask), step)
	for i := 0; i < len(step); i++ {
		step[i] = step[i] & field.fullMask[i]
	}
	field.step = step
}

func (field *LongField) GetCount() uint64 {
	return uint64(field.count)
}

func (field *LongField) SetCount(count uint64) {
	if count > 1 {
		// FIXME
		// field.count = uint64(count % uint64(field.fullMask))
		field.count = count
	} else {
		field.count = uint64(1)
	}
}

func (field *LongField) GetMode() uint8 {
	return field.mode
}

func (field *LongField) SetMode(mode uint8) {
	switch mode {
	case AUTO, FIXED, RANDOMIZE, INCREMENT, DECREMENT:
		field.mode = mode
	default:
		Error.Println("Invalid mode: ", mode, " setting mode to ", FIXED, " (fixed)")
		field.mode = FIXED
	}
}

func (field *LongField) FromCapnp(capnpField *schemas.Field) {
	value, _ := capnpField.Value()
	field.SetValue(value)

	step, _ := capnpField.Step()
	field.SetStep(step)

	mask, _ := capnpField.Mask()
	field.SetMask(mask)

	field.SetMode(capnpField.Mode())
	field.SetCount(capnpField.Count())
}

func (field *LongField) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)
	capnpField.SetValue(field.GetValue())
	capnpField.SetMode(field.GetMode())
	capnpField.SetStep(field.GetStep())
	capnpField.SetMask(field.GetMask())
	capnpField.SetCount(field.GetCount())
	return capnpField
}

func (field *LongField) SetCurrentValue(index uint) {
	if uint64(index)%field.count == 0 {
		field.value = field.firstValue
		return
	}
	switch field.mode {
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
	for i := len(field.value) - 1; i >= 0; i-- {
		field.value[i] = byte(rand.Intn(256))
	}
}

func sub(a, b byte) (result byte, overflow bool) {
	if b > a {
		return 255 - ((b - a) - 1), true
	} else {
		return (a - b), false
	}
}

func (field *LongField) Decrement() {
	overflow := false
	var newValue byte
	for i := len(field.value) - 1; i >= 0; i-- {
		if field.mask[i] == 0 {
			overflow = false
			continue
		}
		if overflow == true {
			newValue, overflow = sub(field.value[i], 1)
		} else {
			newValue = field.value[i]
		}
		if overflow == true {
			newValue, _ = sub(newValue, field.step[i])
		} else {
			newValue, overflow = sub(newValue, field.step[i])
		}
		field.value[i] = newValue & field.mask[i]
	}
}

func add(a, b byte) (result, overflow byte) {
	res := uint16(a) + uint16(b)
	result = byte(res & 255)
	overflow = byte(res >> 8)
	return
}

func (field *LongField) Increment() {
	overflow := byte(0)
	var newValue, tmpOverflow byte
	for i := len(field.value) - 1; i >= 0; i-- {
		if field.mask[i] == 0 {
			overflow = 0
			continue
		}
		if overflow > 0 {
			newValue, tmpOverflow = add(field.value[i], overflow)
		} else {
			newValue = field.value[i]
			tmpOverflow = 0
		}
		newValue, overflow = add(newValue, field.step[i])
		field.value[i] = newValue & field.mask[i]
		overflow += tmpOverflow
	}
}

func toString(field Field) string {
	return fmt.Sprint(
		"{value: ", field.GetValue(),
		", mask: ", field.GetMask(),
		", step: ", field.GetStep(),
		", mode: ", field.GetMode(),
		", count: ", field.GetCount(),
		"}")
	// ", fullMask: ", field.GetFullMask())
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
