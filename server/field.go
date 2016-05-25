package server

import (
	"math/rand"
)

type Field interface {
	increment()
	decrement()
	randomize()
	SetValue()
}

type Field8 struct {
	initValue uint8
	value     uint8
	step      uint8
	count     uint8
	mode      string
	mask      uint8
}

func (field Field8) SetValue(index uint) {
	if uint8(index)%field.count == 0 {
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

func (field Field8) increment() {
	field.value = (field.value & (^field.mask)) | ((field.value + field.step) & field.mask)
}

func (field Field8) decrement() {
	field.value = (field.value & (^field.mask)) | ((field.value - field.step) & field.mask)
}

func (field Field8) randomize() {
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

func (field Field16) SetValue(index uint) {
	if uint16(index)%field.count == 0 {
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

func (field LongField) SetValue(index uint64) {
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

func (field LongField) randomize() {
	for i := len(field.value) - 1; i >= 0; i-- {
		field.value[i] = uint8(rand.Intn(256))
	}
}

func sub(a, b uint8) (result uint8, overflow bool) {
	if b > a {
		return 255 - ((b - a) - 1), true
	} else {
		return (a - b), false
	}
}

func (field LongField) decrement() {
	overflow := false
	var newValue uint8
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

func add(a, b uint8) (result, overflow uint8) {
	res := uint16(a) + uint16(b)
	result = uint8(res & 255)
	overflow = uint8(res >> 8)
	return
}

func (field LongField) increment() {
	overflow := uint8(0)
	var newValue, tmpOverflow uint8
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
