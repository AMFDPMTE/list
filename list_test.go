package main

import "testing"

func TestList_Insert(t *testing.T) {
	list := List{}

	if list.Root != nil {
		t.Fatalf("empty list was expected, got %v", list.Root)
	}

	list.Insert(10)

	if list.Root == nil {
		t.Fatal("list should not be empty")
	}
}

func TestList_Insert_duplicates(t *testing.T) {
	list := List{}
	list.Insert(10)
	list.Insert(10)

	// should be one list in the array
	first := *list.Root

	if first.Next != nil {
		t.Fatalf("Was not expecing the first node to point to anything(nil ptr), got %v", first.Next)
	}
}

func TestLlist_length(t *testing.T) {
	list := List{}
	if e, a := uint(0), list.Length(); e != a {
		t.Fatalf("was expecting a list of length %v, got %v", e, a)
	}

	list.Insert(20)
	if e, a := uint(1), list.Length(); e != a {
		t.Fatalf("was expecting a list of length %v, got %v", e, a)
	}

	list.Insert(10)
	if e, a := uint(2), list.Length(); e != a {
		t.Fatalf("was expecting a list of length %v, got %v", e, a)
	}

	list.Insert(200)
	if e, a := uint(3), list.Length(); e != a {
		t.Fatalf("was expecting a list of length %v, got %v", e, a)
	}
}

func TestList_Contains(t *testing.T) {
	list := List{}

	if e, a := false, list.Contain(2); e != a {
		t.Fatalf("was expecting to return %v, but return %v", e, a)
	}

	list.Insert(20)
	list.Insert(5)
	list.Insert(15)

	if e, a := true, list.Contain(15); e != a {
		t.Fatalf("was expecting to return %v, but return %v", e, a)
	}

	if e, a := false, list.Contain(2); e != a {
		t.Fatalf("was expecting to return %v, but return %v", e, a)
	}
}
