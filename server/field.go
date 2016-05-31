package server

import (
	schemas "github.com/little-dude/tgen/capnp"
	"math/rand"
	"zombiezen.com/go/capnproto2"
)

type Field interface {
	increment()
	decrement()
	randomize()
	SetValue(uint)
	ToCapnp(*capnp.Segment) schemas.Field
	FromCapnp(*schemas.Field)
}

type Field8 struct {
	initValue uint8
	value     uint8
	step      uint8
	count     uint8
	mode      string
	mask      uint8
}

func NewField(field schemas.Field) Field {
	value, _ := field.Value()
	length := len(value)
	if length == 1 {
		return &Field8{}
	}
	if length == 2 {
		return &Field16{}
	}
	if length <= 4 {
		return &Field32{}
	}
	if length <= 8 {
		return &Field64{}
	}
	return &LongField{}
}

func (field *Field8) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)
	capnpField.SetValue([]byte{byte(field.value)})
	capnpField.SetMode(field.mode)
	capnpField.SetStep([]byte{byte(field.step)})
	capnpField.SetMask([]byte{byte(field.mask)})
	capnpField.SetCount(uint64(field.count))
	return capnpField
}

func (field *Field8) FromCapnp(capnpField *schemas.Field) {
	var bytes []byte
	bytes, _ = capnpField.Value()
	field.value = uint8(bytes[0])

	bytes, _ = capnpField.Step()
	field.step = uint8(bytes[0])

	bytes, _ = capnpField.Mask()
	field.mask = uint8(bytes[0])

	field.mode, _ = capnpField.Mode()

	field.count = uint8(capnpField.Count())
}

func (field *Field8) SetValue(index uint) {
	if index%uint(field.count) == 0 {
		field.value = field.initValue
		return
	}
	switch field.mode {
	case "increment":
		field.increment()
	case "decrement":
		field.decrement()
	case "random":
		field.randomize()
	}
}

func (field *Field8) increment() {
	field.value = (field.value & (^field.mask)) | ((field.value + field.step) & field.mask)
}

func (field *Field8) decrement() {
	field.value = (field.value & (^field.mask)) | ((field.value - field.step) & field.mask)
}

func (field *Field8) randomize() {
	field.value = uint8(rand.Int31()) & field.mask
}

type Field16 struct {
	initValue  uint16
	resetIndex uint16
	value      uint16
	step       uint16
	count      uint16
	mode       string
	mask       uint16
}

func (field *Field16) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)
	capnpField.SetValue([]byte{
		byte(field.value >> 8),
		byte(field.value & 255),
	})
	capnpField.SetMode(field.mode)
	capnpField.SetStep([]byte{
		byte(field.step >> 8),
		byte(field.step & 255),
	})
	capnpField.SetMask([]byte{
		byte(field.mask >> 8),
		byte(field.mask & 255),
	})
	capnpField.SetCount(uint64(field.count))
	return capnpField
}

func (field *Field16) FromCapnp(capnpField *schemas.Field) {
	var bytes []byte
	bytes, _ = capnpField.Value()
	field.value = uint16(bytes[0]) + (uint16(bytes[1]) << 8)

	bytes, _ = capnpField.Step()
	field.step = uint16(bytes[0]) + (uint16(bytes[1]) << 8)

	bytes, _ = capnpField.Mask()
	field.mask = uint16(bytes[0]) + (uint16(bytes[1]) << 8)

	field.mode, _ = capnpField.Mode()

	field.count = uint16(capnpField.Count() % uint64(^uint16(0)))
}

func (field Field16) SetValue(index uint) {
	if index%uint(field.count) == 0 {
		field.value = field.initValue
		return
	}
	switch field.mode {
	case "increment":
		field.increment()
	case "decrement":
		field.decrement()
	case "random":
		field.randomize()
	}
}

func (field Field16) increment() {
	field.value = (field.value & (^field.mask)) | ((field.value + field.step) & field.mask)
}

func (field Field16) decrement() {
	field.value = (field.value & (^field.mask)) | ((field.value - field.step) & field.mask)
}

func (field Field16) randomize() {
	field.value = uint16(rand.Int31()) & field.mask
}

type Field32 struct {
	initValue  uint32
	resetIndex uint32
	value      uint32
	step       uint32
	count      uint32
	mode       string
	mask       uint32
}

func (field *Field32) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)
	value := make([]byte, 4)
	step := make([]byte, 4)
	mask := make([]byte, 4)
	for i := 0; i < 4; i++ {
		value[i] = byte((field.value >> uint(8*(3-i))) & 255)
		mask[i] = byte((field.mask >> uint(8*(3-i))) & 255)
		step[i] = byte((field.step >> uint(8*(3-i))) & 255)
	}
	capnpField.SetValue(value)
	capnpField.SetStep(step)
	capnpField.SetMask(mask)
	capnpField.SetMode(field.mode)
	capnpField.SetCount(uint64(field.count))
	return capnpField
}

func (field *Field32) FromCapnp(capnpField *schemas.Field) {
	value, _ := capnpField.Value()
	mask, _ := capnpField.Mask()
	step, _ := capnpField.Step()
	for i := 0; i < 4; i++ {
		field.value += uint32(value[i]) << uint(3-i)
		field.mask += uint32(mask[i]) << uint(3-i)
		field.step += uint32(step[i]) << uint(3-i)
	}
	field.mode, _ = capnpField.Mode()
	field.count = uint32(capnpField.Count() % uint64(^uint32(0)))
}

func (field Field32) SetValue(index uint) {
	if uint32(index)%field.count == 0 {
		field.value = field.initValue
		return
	}
	switch field.mode {
	case "increment":
		field.increment()
	case "decrement":
		field.decrement()
	case "random":
		field.randomize()
	}
}

func (field Field32) increment() {
	field.value = (field.value & (^field.mask)) | ((field.value + field.step) & field.mask)
}

func (field Field32) decrement() {
	field.value = (field.value & (^field.mask)) | ((field.value - field.step) & field.mask)
}

func (field Field32) randomize() {
	field.value = uint32(rand.Int31()) & field.mask
}

type Field64 struct {
	initValue  uint64
	resetIndex uint64
	value      uint64
	step       uint64
	count      uint64
	mode       string
	mask       uint64
}

func (field *Field64) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)
	value := make([]byte, 8)
	step := make([]byte, 8)
	mask := make([]byte, 8)
	for i := 0; i < 8; i++ {
		value[i] = byte((field.value >> uint(8*(3-i))) & 255)
		mask[i] = byte((field.mask >> uint(8*(3-i))) & 255)
		step[i] = byte((field.step >> uint(8*(3-i))) & 255)
	}
	capnpField.SetValue(value)
	capnpField.SetStep(step)
	capnpField.SetMask(mask)
	capnpField.SetMode(field.mode)
	capnpField.SetCount(field.count)
	return capnpField
}

func (field *Field64) FromCapnp(capnpField *schemas.Field) {
	value, _ := capnpField.Value()
	mask, _ := capnpField.Mask()
	step, _ := capnpField.Step()
	for i := 0; i < 8; i++ {
		field.value += uint64(value[i]) << uint(3-i)
		field.mask += uint64(mask[i]) << uint(3-i)
		field.step += uint64(step[i]) << uint(3-i)
	}
	field.mode, _ = capnpField.Mode()
	field.count = capnpField.Count()
}

func (field Field64) SetValue(index uint) {
	if uint64(index)%field.count == 0 {
		field.value = field.initValue
		return
	}
	switch field.mode {
	case "increment":
		field.increment()
	case "decrement":
		field.decrement()
	case "random":
		field.randomize()
	}
}

func (field Field64) increment() {
	field.value = (field.value & (^field.mask)) | ((field.value + field.step) & field.mask)
}

func (field Field64) decrement() {
	field.value = (field.value & (^field.mask)) | ((field.value - field.step) & field.mask)
}

func (field Field64) randomize() {
	field.value = uint64(rand.Int63()) & field.mask
}

type LongField struct {
	initValue []byte
	value     []byte
	step      []byte
	count     uint64 // does not really make sense to generate more than 2^64 different packets
	mode      string
	mask      []byte
}

func (field *LongField) ToCapnp(seg *capnp.Segment) (capnpField schemas.Field) {
	capnpField, _ = schemas.NewField(seg)

	tmpSlice := make([]byte, len(field.value))
	copy(tmpSlice, field.value)
	capnpField.SetValue(tmpSlice)

	tmpSlice = make([]byte, len(field.mask))
	copy(tmpSlice, field.mask)
	capnpField.SetMask(tmpSlice)

	tmpSlice = make([]byte, len(field.step))
	copy(tmpSlice, field.step)
	capnpField.SetStep(tmpSlice)

	capnpField.SetMode(field.mode)
	capnpField.SetCount(field.count)
	return capnpField
}

func (field *LongField) FromCapnp(capnpField *schemas.Field) {
	field.value, _ = capnpField.Value()
	field.step, _ = capnpField.Step()
	field.mask, _ = capnpField.Mask()
	field.mode, _ = capnpField.Mode()
	field.count = capnpField.Count()
}

func (field *LongField) SetValue(index uint) {
	if uint64(index)%field.count == 0 {
		field.value = field.initValue
		return
	}
	switch field.mode {
	case "increment":
		field.increment()
	case "decrement":
		field.decrement()
	case "randomize":
		field.randomize()
	default:
		return
	}
}

func (field *LongField) randomize() {
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

func (field *LongField) decrement() {
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

func (field *LongField) increment() {
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
