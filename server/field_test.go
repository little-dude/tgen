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
		field:    Field8{initValue: 0, value: 0, step: 0, count: 1, mode: "increment", mask: 0},
		index:    []uint{0, 1},
		expected: []uint8{0, 0},
	},
	{
		field:    Field8{initValue: 0, value: 0, step: 1, count: 1, mode: "increment", mask: 255},
		index:    []uint{0, 1},
		expected: []uint8{0, 0},
	},
	{
		field:    Field8{initValue: 0, value: 0, step: 0, count: 2, mode: "increment", mask: 255},
		index:    []uint{0, 1, 255},
		expected: []uint8{0, 0, 0},
	},
	{
		field:    Field8{initValue: 0, value: 0, step: 1, count: 2, mode: "increment", mask: 255},
		index:    []uint{0, 1, 2},
		expected: []uint8{0, 1, 0},
	},
	{
		field:    Field8{initValue: 0, value: 0, step: 1, count: 255, mode: "increment", mask: 255},
		index:    []uint{255, 256},
		expected: []uint8{0, 1},
	},
	{
		field:    Field8{initValue: 255, value: 255, step: 1, count: 255, mode: "increment", mask: 0},
		index:    []uint{0, 1, 2},
		expected: []uint8{255, 255, 255},
	},
	{
		field:    Field8{initValue: 0, value: 0, step: 255, count: 255, mode: "increment", mask: 64},
		index:    []uint{0, 1, 2},
		expected: []uint8{0, 64, 0},
	},
	{
		field:    Field8{initValue: 0, value: 0, step: 32, count: 255, mode: "increment", mask: 255},
		index:    []uint{0, 1, 2, 3, 4, 5, 6, 7, 8},
		expected: []uint8{0, 32, 64, 96, 128, 160, 192, 224, 0},
	},
}

func TestField8(t *testing.T) {
	var values []uint8
	for _, test := range tests8 {
		values = make([]uint8, len(test.expected))
		for i, index := range test.index {
			test.field.SetValue(index)
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

type testLongField struct {
	field    LongField
	index    []uint
	expected [][]byte
}

var testsLong = []testLongField{
	{
		field: LongField{
			initValue: []byte{0, 0, 0, 0, 0, 0},
			value:     []byte{0, 0, 0, 0, 0, 0},
			step:      []byte{0, 0, 0, 0, 0, 0},
			count:     1,
			mode:      "increment",
			mask:      []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{0, 1},
		expected: [][]byte{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}},
	},
	{
		field: LongField{
			initValue: []byte{0, 0, 0, 0, 0, 0},
			value:     []byte{0, 0, 0, 0, 0, 0},
			step:      []byte{1, 1, 1, 1, 1, 1},
			count:     2,
			mode:      "increment",
			mask:      []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{0, 1},
		expected: [][]byte{{0, 0, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1}},
	},
	{
		field: LongField{
			initValue: []byte{0, 0, 0, 0, 0, 0},
			value:     []byte{0, 0, 0, 255, 0, 0},
			step:      []byte{0, 0, 0, 1, 0, 0},
			count:     2,
			mode:      "increment",
			mask:      []byte{255, 255, 255, 255, 255, 255},
		},
		index:    []uint{1},
		expected: [][]byte{{0, 0, 1, 0, 0, 0}},
	},
	{
		field: LongField{
			initValue: []byte{0, 0, 0, 0, 0, 0},
			value:     []byte{0, 0, 0, 255, 0, 0},
			step:      []byte{0, 0, 0, 1, 0, 0},
			count:     2,
			mode:      "increment",
			mask:      []byte{255, 255, 0, 255, 255, 255},
		},
		index:    []uint{1},
		expected: [][]byte{{0, 0, 0, 0, 0, 0}},
	},
	{
		field: LongField{
			initValue: []byte{1, 2, 3, 4, 5, 6},
			value:     []byte{255, 255, 255, 255, 255, 255},
			step:      []byte{0, 0, 0, 0, 0, 255},
			count:     1000,
			mode:      "increment",
			mask:      []byte{255, 255, 255, 255, 255, 255},
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
			test.field.SetValue(index)
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
