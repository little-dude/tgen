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
		field:    Field8{FirstValue: 0, Value: 0, Step: 0, Count: 1, Mode: INCREMENT, Mask: 0, FullMask: 255},
		index:    []uint{0, 1},
		expected: []uint8{0, 0},
	},
	{
		field:    Field8{FirstValue: 0, Value: 0, Step: 1, Count: 1, Mode: INCREMENT, Mask: 255, FullMask: 255},
		index:    []uint{0, 1},
		expected: []uint8{0, 0},
	},
	{
		field:    Field8{FirstValue: 0, Value: 0, Step: 0, Count: 2, Mode: INCREMENT, Mask: 255, FullMask: 255},
		index:    []uint{0, 1, 255},
		expected: []uint8{0, 0, 0},
	},
	{
		field:    Field8{FirstValue: 0, Value: 0, Step: 1, Count: 2, Mode: INCREMENT, Mask: 255, FullMask: 255},
		index:    []uint{0, 1, 2},
		expected: []uint8{0, 1, 0},
	},
	{
		field:    Field8{FirstValue: 0, Value: 0, Step: 1, Count: 255, Mode: INCREMENT, Mask: 255, FullMask: 255},
		index:    []uint{255, 256},
		expected: []uint8{0, 1},
	},
	{
		field:    Field8{FirstValue: 255, Value: 255, Step: 1, Count: 255, Mode: INCREMENT, Mask: 0, FullMask: 255},
		index:    []uint{0, 1, 2},
		expected: []uint8{255, 255, 255},
	},
	{
		field:    Field8{FirstValue: 0, Value: 0, Step: 255, Count: 255, Mode: INCREMENT, Mask: 64, FullMask: 255},
		index:    []uint{0, 1, 2},
		expected: []uint8{0, 64, 0},
	},
	{
		field:    Field8{FirstValue: 0, Value: 0, Step: 32, Count: 255, Mode: INCREMENT, Mask: 255, FullMask: 255},
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
			values[i] = test.field.Value
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
	field = &Field8{FirstValue: 0, Value: 0, Step: 32, Count: 255, Mode: INCREMENT, Mask: 255, FullMask: 255}
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
	field = &Field8{FullMask: 255}

	field.SetValue([]byte{255})
	if !reflect.DeepEqual(field.Value, uint8(255)) {
		t.Error(field.Value)
	}

	field.SetMask([]byte{32})
	if !reflect.DeepEqual(field.Mask, uint8(32)) {
		t.Error(field.Mask)
	}

	field.SetStep([]byte{0})
	if !reflect.DeepEqual(field.Step, uint8(0)) {
		t.Error(field.Step)
	}
}

func TestField16Getters(t *testing.T) {
	var field *Field16

	field = &Field16{FirstValue: 0, Value: 0, Step: 32, Count: 255, Mode: INCREMENT, Mask: 255, FullMask: 65535}
	if !reflect.DeepEqual(field.GetMask(), []byte{0, 255}) {
		t.Error(field.GetMask())
	}
	if !reflect.DeepEqual(field.GetStep(), []byte{0, 32}) {
		t.Error(field.GetStep())
	}
	if !reflect.DeepEqual(field.GetValue(), []byte{0, 0}) {
		t.Error(field.GetValue())
	}

	field = &Field16{FirstValue: 0, Value: 65535, Step: 256, Count: 1025, Mode: INCREMENT, Mask: 1025, FullMask: 65535}
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
	field = &Field16{FullMask: 65535}

	field.SetValue([]byte{0, 255})
	if !reflect.DeepEqual(field.Value, uint16(255)) {
		t.Error(field.Value)
	}
	field.SetValue([]byte{255, 255})
	if !reflect.DeepEqual(field.Value, uint16(65535)) {
		t.Error(field.Value)
	}

	field.SetMask([]byte{0, 32})
	if !reflect.DeepEqual(field.Mask, uint16(32)) {
		t.Error(field.Mask)
	}
	field.SetMask([]byte{1, 0})
	if !reflect.DeepEqual(field.Mask, uint16(256)) {
		t.Error(field.Mask)
	}

	field.SetStep([]byte{0, 0})
	if !reflect.DeepEqual(field.Step, uint16(0)) {
		t.Error(field.Step)
	}
	field.SetStep([]byte{4, 1})
	if !reflect.DeepEqual(field.Step, uint16(1025)) {
		t.Error(field.Step)
	}

	field = &Field16{FullMask: 0x01ff}

	field.SetValue([]byte{255, 255})
	if !reflect.DeepEqual(field.Value, uint16(0x01ff)) {
		t.Error(field.Value)
	}

	field.SetMask([]byte{1, 0})
	if !reflect.DeepEqual(field.Mask, uint16(0x0100)) {
		t.Error(field.Mask)
	}
	field.SetStep([]byte{4, 1})
	if !reflect.DeepEqual(field.Step, uint16(0x0001)) {
		t.Error(field.Step)
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
			FirstValue: []byte{0, 0, 0, 0, 0, 0},
			Value:      []byte{0, 0, 0, 0, 0, 0},
			Step:       []byte{0, 0, 0, 0, 0, 0},
			Count:      1,
			Mode:       INCREMENT,
			Mask:       []byte{255, 255, 255, 255, 255, 255},
			FullMask:   []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{0, 1},
		expected: [][]byte{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}},
	},
	{
		field: LongField{
			FirstValue: []byte{0, 0, 0, 0, 0, 0},
			Value:      []byte{0, 0, 0, 0, 0, 0},
			Step:       []byte{1, 1, 1, 1, 1, 1},
			Count:      2,
			Mode:       INCREMENT,
			Mask:       []byte{255, 255, 255, 255, 255, 255},
			FullMask:   []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{0, 1},
		expected: [][]byte{{0, 0, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1}},
	},
	{
		field: LongField{
			FirstValue: []byte{0, 0, 0, 0, 0, 0},
			Value:      []byte{0, 0, 0, 255, 0, 0},
			Step:       []byte{0, 0, 0, 1, 0, 0},
			Count:      2,
			Mode:       INCREMENT,
			Mask:       []byte{255, 255, 255, 255, 255, 255},
			FullMask:   []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{1},
		expected: [][]byte{{0, 0, 1, 0, 0, 0}},
	},
	{
		field: LongField{
			FirstValue: []byte{0, 0, 0, 0, 0, 0},
			Value:      []byte{0, 0, 0, 255, 0, 0},
			Step:       []byte{0, 0, 0, 1, 0, 0},
			Count:      2,
			Mode:       INCREMENT,
			Mask:       []byte{255, 255, 0, 255, 255, 255},
			FullMask:   []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{1},
		expected: [][]byte{{0, 0, 0, 0, 0, 0}},
	},
	{
		field: LongField{
			FirstValue: []byte{1, 2, 3, 4, 5, 6},
			Value:      []byte{255, 255, 255, 255, 255, 255},
			Step:       []byte{0, 0, 0, 0, 0, 255},
			Count:      1000,
			Mode:       INCREMENT,
			Mask:       []byte{255, 255, 255, 255, 255, 255},
			FullMask:   []byte{255, 255, 255, 255, 255, 255},
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
			values[i] = make([]byte, len(test.field.Value))
			copy(values[i], test.field.Value)
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
