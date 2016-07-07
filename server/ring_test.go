package server

import (
	"reflect"
	"testing"
)

var b1 = Buffer{&RawPacket{}}
var b2 = Buffer{&RawPacket{}}
var b3 = Buffer{&RawPacket{}}
var b4 = Buffer{&RawPacket{}}
var b5 = Buffer{&RawPacket{}}
var b6 = Buffer{&RawPacket{}}
var b7 = Buffer{&RawPacket{}}
var b8 = Buffer{&RawPacket{}}
var b9 = Buffer{&RawPacket{}}
var b10 = Buffer{&RawPacket{}}
var buffers = [...]Buffer{b1, b2, b3, b4, b5, b6, b7, b8, b9, b10}

func TestSet(t *testing.T) {
	ring := newRingBuf(10)
	data := []Buffer{b1, b2, b3, b4, b5, b6, b7, b8, b9, b10}
	for i := 0; i < 10; i++ {
		ring.set(data[i])
		if ring.tail != 0 {
			t.Fatal("no consumer, tail should remain 0, but is", ring.tail)
		}
		if ring.head != i {
			t.Fatal("unexpected index", ring.head, "for head")
		}
	}
	if !reflect.DeepEqual(ring.buff, data) {
		t.Fatal(ring.buff, "!=", data)
	}
}

func TestGet(t *testing.T) {
	ring := newRingBuf(10)
	data := []Buffer{b1, b2, b3, b4, b5, b6, b7, b8, b9, b10}
	ring.buff = data
	ring.tail = 0
	ring.head = 9

	for i := 0; i < 10; i++ {
		b := ring.get()
		if !reflect.DeepEqual(b, data[i]) {
			t.Fatal(b, "!=", data[i])
		}
		if ring.tail != i+1 {
			t.Fatal("unexpected index", ring.tail, "for tail")
		}
		if ring.head != 9 {
			t.Fatal("no consumer, header should remain 9 but is", ring.head)
		}
		if !reflect.DeepEqual(ring.buff, data) {
			t.Fatal(ring.buff, "!=", data)
		}
	}
}

func TestShrink(t *testing.T) {

	data := make([]Buffer, 0)
	for i := 0; i < 10; i++ {
		data = append(data, []Buffer{b1, b2, b3, b4, b5, b6, b7, b8, b9, b10}...)
	}

	// Test 1: shrink when r.tail < r.head
	ring := newRingBuf(10)
	ring.len = 100
	ring.buff = make([]Buffer, 100)
	copy(ring.buff, data)
	ring.tail = 30 // ring.buff[ring.tail] = b1
	ring.head = 35 // ring.buff[ring.tail] = b6

	ring.resize(10)
	if ring.len != 10 {
		t.Fatal("len should be 10 not", ring.len)
	}
	if ring.tail != 0 {
		t.Fatal("tail should be 0 not", ring.tail)
	}
	if ring.head != 5 {
		t.Fatal("head should be 5 not", ring.head)
	}
	if !reflect.DeepEqual(ring.buff[0:6], []Buffer{b1, b2, b3, b4, b5, b6}) {
		t.Fatal("wrong buffer values:", ring.buff[0:6])
	}

	// Test 2: shrink when r.head % r.len < r.tail % r.len
	ring = newRingBuf(10)
	ring.len = 100
	ring.buff = make([]Buffer, 100)
	copy(ring.buff, data)
	ring.tail = 298 // ring.buff[ring.tail] = b9
	ring.head = 305 // ring.buff[ring.head] = b6

	ring.resize(10)
	if ring.len != 10 {
		t.Fatal("len should be 10 not", ring.len)
	}
	if ring.tail != 0 {
		t.Fatal("tail should be 0 not", ring.tail)
	}
	if ring.head != 7 {
		t.Fatal("head should be 7 not", ring.head)
	}
	if !reflect.DeepEqual(ring.buff[0:8], []Buffer{b9, b10, b1, b2, b3, b4, b5, b6}) {
		t.Fatal("wrong buffer values:", ring.buff[0:8])
	}
}

func TestExtend(t *testing.T) {
	data := []Buffer{b1, b2, b3, b4, b5, b6, b7, b8, b9, b10}
	ring := newRingBuf(10)
	copy(ring.buff, data)

	ring.tail = 7
	ring.head = 16
	ring.resize(40)
	if ring.len != 40 {
		t.Fatal("len should be 40 not", ring.len)
	}
	if ring.tail != 0 {
		t.Fatal("tail should be 0 not", ring.tail)
	}
	if ring.head != 9 {
		t.Fatal("head should be 9 not", ring.head)
	}
	expected := []Buffer{b8, b9, b10, b1, b2, b3, b4, b5, b6, b7}
	if !reflect.DeepEqual(ring.buff[0:10], expected) {
		t.Fatal("wrong buffer values:", ring.buff[0:10], "!=", expected)
	}
}
