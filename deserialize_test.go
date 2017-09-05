package main

import "testing"

func TestDeserialize(t *testing.T) {
	data := []byte{
		0x00, 0x05, // 5
		0x00, 0x0A, // 10
		0x00, 0x0F, // 15
	}
	list, err := Deserialize(data)

	expected := []uint16{5, 10, 15}

	if err != nil {
		t.Fatalf("Was not expecting an error but got one: %s", err)
	}

	if e, a := uint(3), list.Length(); e != a {
		t.Fatalf("Was expecting a list of length %v, got %v", e, a)
	}

	for _, e := range expected {
		if !list.Contain(e) {
			t.Fatalf("Was expecting the list to contain %v, but it did not", e)
		}
	}
}

func TestDeserialize_empty(t *testing.T) {
	data := []byte{}
	list, err := Deserialize(data)

	if err != nil {
		t.Fatalf("Was not expecting an error but did not get one")
	}

	if e, a := uint(0), list.Length(); e != a {
		t.Fatalf("Was expecting a list of length %v, got %v", e, a)
	}
}

func TestDeserialize_invalidBlob(t *testing.T) {
	data := []byte{
		0x00, 0x05, // 5
		0xFF, // only one byte, we are storing 16 bit integers
	}

	_, err := Deserialize(data)
	if err == nil {
		t.Fatalf("Was expecting an error but did not get one")
	}
}

func TestDeserialize_unorderedBlob(t *testing.T) {
	data := []byte{
		0x00, 0x05, // 5
		0x00, 0x0F, // 15
		0x00, 0x0A, // 10
	} // this blob is invalid as it was out of numerical order

	_, err := Deserialize(data)
	if err == nil {
		t.Fatalf("Was expecting an error but did not get one")
	}
}
