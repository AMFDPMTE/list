package main

import "testing"

func TestList_Serialize(t *testing.T) {
	list := List{}
	list.Insert(5)
	list.Insert(10)
	list.Insert(15)

	data := list.Serialize()

	expected := []byte{
		0x00,
		0x05,
		0x00,
		0x0A,
		0x00,
		0x0F,
	}

	if e, a := len(expected), len(data); e != a {
		t.Fatalf("Was expecting a byte slice of len %v, got %v", e, a)
	}

	for index, e := range expected {
		if a := data[index]; e != a {
			t.Fatalf("Was expecting %v at index %v, got %v", e, index, a)
		}
	}
}

func Test_Serialize_empty(t *testing.T) {
	list := List{}
	data := list.Serialize()

	expected := []byte{}

	if e, a := len(expected), len(data); e != a {
		t.Fatalf("Was expecting a byte slice of len %v, got %v", e, a)
	}
}
