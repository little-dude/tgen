package server

import (
	"reflect"
	"testing"
)

type testField8 struct {
	field    Field8
	index    []uint
	expected []uint8
}

var tests8 = []testField8{
	{
		field:    Field8{firstValue: 0, value: 0, step: 0, count: 1, mode: INCREMENT, mask: 0, fullMask: 255},
		index:    []uint{0, 1},
		expected: []uint8{0, 0},
	},
	{
		field:    Field8{firstValue: 0, value: 0, step: 1, count: 1, mode: INCREMENT, mask: 255, fullMask: 255},
		index:    []uint{0, 1},
		expected: []uint8{0, 0},
	},
	{
		field:    Field8{firstValue: 0, value: 0, step: 0, count: 2, mode: INCREMENT, mask: 255, fullMask: 255},
		index:    []uint{0, 1, 255},
		expected: []uint8{0, 0, 0},
	},
	{
		field:    Field8{firstValue: 0, value: 0, step: 1, count: 2, mode: INCREMENT, mask: 255, fullMask: 255},
		index:    []uint{0, 1, 2},
		expected: []uint8{0, 1, 0},
	},
	{
		field:    Field8{firstValue: 0, value: 0, step: 1, count: 255, mode: INCREMENT, mask: 255, fullMask: 255},
		index:    []uint{255, 256},
		expected: []uint8{0, 1},
	},
	{
		field:    Field8{firstValue: 255, value: 255, step: 1, count: 255, mode: INCREMENT, mask: 0, fullMask: 255},
		index:    []uint{0, 1, 2},
		expected: []uint8{255, 255, 255},
	},
	{
		field:    Field8{firstValue: 0, value: 0, step: 255, count: 255, mode: INCREMENT, mask: 64, fullMask: 255},
		index:    []uint{0, 1, 2},
		expected: []uint8{0, 64, 0},
	},
	{
		field:    Field8{firstValue: 0, value: 0, step: 32, count: 255, mode: INCREMENT, mask: 255, fullMask: 255},
		index:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8},
		expected: []uint8{0, 32, 64, 96, 128, 160, 192, 224, 0},
	},
}

func TestField8Increment(t *testing.T) {
	var values []uint8
	for _, test := range tests8 {
		values = make([]uint8, len(test.expected))
		for i, index := range test.index {
			test.field.SetCurrentValue(index)
			values[i] = test.field.value
		}
		if reflect.DeepEqual(values, test.expected) == false {
			t.Error(
				"For ", test,
				"expected ", test.expected,
				"got ", values,
			)
		}
	}
}

func TestField8Getters(t *testing.T) {
	var field *Field8
	field = &Field8{firstValue: 0, value: 0, step: 32, count: 255, mode: INCREMENT, mask: 255, fullMask: 255}
	if !reflect.DeepEqual(field.GetMask(), []byte{255}) {
		t.Error(field.GetMask())
	}
	if !reflect.DeepEqual(field.GetStep(), []byte{32}) {
		t.Error(field.GetStep())
	}
	if !reflect.DeepEqual(field.GetValue(), []byte{0}) {
		t.Error(field.GetValue())
	}
}

func TestField8Setters(t *testing.T) {
	var field *Field8
	field = &Field8{fullMask: 255}

	field.SetValue([]byte{255})
	if !reflect.DeepEqual(field.value, uint8(255)) {
		t.Error(field.value)
	}

	field.SetMask([]byte{32})
	if !reflect.DeepEqual(field.mask, uint8(32)) {
		t.Error(field.mask)
	}

	field.SetStep([]byte{0})
	if !reflect.DeepEqual(field.step, uint8(0)) {
		t.Error(field.step)
	}
}

func TestField16Getters(t *testing.T) {
	var field *Field16

	field = &Field16{firstValue: 0, value: 0, step: 32, count: 255, mode: INCREMENT, mask: 255, fullMask: 65535}
	if !reflect.DeepEqual(field.GetMask(), []byte{0, 255}) {
		t.Error(field.GetMask())
	}
	if !reflect.DeepEqual(field.GetStep(), []byte{0, 32}) {
		t.Error(field.GetStep())
	}
	if !reflect.DeepEqual(field.GetValue(), []byte{0, 0}) {
		t.Error(field.GetValue())
	}

	field = &Field16{firstValue: 0, value: 65535, step: 256, count: 1025, mode: INCREMENT, mask: 1025, fullMask: 65535}
	if !reflect.DeepEqual(field.GetMask(), []byte{4, 1}) {
		t.Error(field.GetMask())
	}
	if !reflect.DeepEqual(field.GetStep(), []byte{1, 0}) {
		t.Error(field.GetStep())
	}
	if !reflect.DeepEqual(field.GetValue(), []byte{255, 255}) {
		t.Error(field.GetValue())
	}
}

func TestField16Setters(t *testing.T) {
	var field *Field16
	field = &Field16{fullMask: 65535}

	field.SetValue([]byte{0, 255})
	if !reflect.DeepEqual(field.value, uint16(255)) {
		t.Error(field.value)
	}
	field.SetValue([]byte{255, 255})
	if !reflect.DeepEqual(field.value, uint16(65535)) {
		t.Error(field.value)
	}

	field.SetMask([]byte{0, 32})
	if !reflect.DeepEqual(field.mask, uint16(32)) {
		t.Error(field.mask)
	}
	field.SetMask([]byte{1, 0})
	if !reflect.DeepEqual(field.mask, uint16(256)) {
		t.Error(field.mask)
	}

	field.SetStep([]byte{0, 0})
	if !reflect.DeepEqual(field.step, uint16(0)) {
		t.Error(field.step)
	}
	field.SetStep([]byte{4, 1})
	if !reflect.DeepEqual(field.step, uint16(1025)) {
		t.Error(field.step)
	}

	field = &Field16{fullMask: 0x01ff}

	field.SetValue([]byte{255, 255})
	if !reflect.DeepEqual(field.value, uint16(0x01ff)) {
		t.Error(field.value)
	}

	field.SetMask([]byte{1, 0})
	if !reflect.DeepEqual(field.mask, uint16(0x0100)) {
		t.Error(field.mask)
	}
	field.SetStep([]byte{4, 1})
	if !reflect.DeepEqual(field.step, uint16(0x0001)) {
		t.Error(field.step)
	}
}

type testLongField struct {
	field    LongField
	index    []uint
	expected [][]byte
}

var testsLong = []testLongField{
	{
		field: LongField{
			firstValue: []byte{0, 0, 0, 0, 0, 0},
			value:      []byte{0, 0, 0, 0, 0, 0},
			step:       []byte{0, 0, 0, 0, 0, 0},
			count:      1,
			mode:       INCREMENT,
			mask:       []byte{255, 255, 255, 255, 255, 255},
			fullMask:   []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{0, 1},
		expected: [][]byte{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}},
	},
	{
		field: LongField{
			firstValue: []byte{0, 0, 0, 0, 0, 0},
			value:      []byte{0, 0, 0, 0, 0, 0},
			step:       []byte{1, 1, 1, 1, 1, 1},
			count:      2,
			mode:       INCREMENT,
			mask:       []byte{255, 255, 255, 255, 255, 255},
			fullMask:   []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{0, 1},
		expected: [][]byte{{0, 0, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1}},
	},
	{
		field: LongField{
			firstValue: []byte{0, 0, 0, 0, 0, 0},
			value:      []byte{0, 0, 0, 255, 0, 0},
			step:       []byte{0, 0, 0, 1, 0, 0},
			count:      2,
			mode:       INCREMENT,
			mask:       []byte{255, 255, 255, 255, 255, 255},
			fullMask:   []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{1},
		expected: [][]byte{{0, 0, 1, 0, 0, 0}},
	},
	{
		field: LongField{
			firstValue: []byte{0, 0, 0, 0, 0, 0},
			value:      []byte{0, 0, 0, 255, 0, 0},
			step:       []byte{0, 0, 0, 1, 0, 0},
			count:      2,
			mode:       INCREMENT,
			mask:       []byte{255, 255, 0, 255, 255, 255},
			fullMask:   []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{1},
		expected: [][]byte{{0, 0, 0, 0, 0, 0}},
	},
	{
		field: LongField{
			firstValue: []byte{1, 2, 3, 4, 5, 6},
			value:      []byte{255, 255, 255, 255, 255, 255},
			step:       []byte{0, 0, 0, 0, 0, 255},
			count:      1000,
			mode:       INCREMENT,
			mask:       []byte{255, 255, 255, 255, 255, 255},
			fullMask:   []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{999, 1000},
		expected: [][]byte{{0, 0, 0, 0, 0, 254}, {1, 2, 3, 4, 5, 6}},
	},
}

func TestLongField(t *testing.T) {
	var values [][]byte
	for _, test := range testsLong {
		values = make([][]byte, len(test.expected))
		for i, index := range test.index {
			test.field.SetCurrentValue(index)
			values[i] = make([]byte, len(test.field.value))
			copy(values[i], test.field.value)
		}
		if reflect.DeepEqual(values, test.expected) == false {
			t.Error(
				"For ", test,
				"expected ", test.expected,
				"got ", values,
			)
		}
	}
}
