package main

import (
	"testing"
)

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

func TestList_Length(t *testing.T) {
	list := List{}

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

func TestList_Length_empty(t *testing.T) {
	list := List{}
	if e, a := uint(0), list.Length(); e != a {
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

func TestList_Remove_nil_List(t *testing.T) {
	list := List{}
	err := list.Remove(22)

	if err == nil {
		t.Fatal("Was expecting an error, but did not get one")
	}
}

func TestList_remove_first_List_many(t *testing.T) {
	list := List{}

	list.Insert(5)
	list.Insert(10)
	list.Insert(15)
	list.Insert(20)
	list.Remove(5)

	if e, a := false, list.Contain(5); e != a {
		t.Fatalf("was expecting to return %v, but return %v", e, a)
	}
}

func TestList_remove_first_List_one(t *testing.T) {
	list := List{}

	list.Insert(5)
	list.Remove(5)

	if e, a := false, list.Contain(5); e != a {
		t.Fatalf("was expecting to return %v, but return %v", e, a)
	}
}

func TestList_remove_last_list_five(t *testing.T) {
	list := List{}

	list.Insert(5)
	list.Insert(10)
	list.Insert(15)
	list.Insert(20)
	list.Insert(25)
	list.Remove(25)

	if e, a := false, list.Contain(25); e != a {
		t.Fatalf("was expecting to return %v, but return %v", e, a)
	}
}

func TestList_remove_remove_all_ten(t *testing.T) {
	list := List{}

	list.Insert(5)
	list.Insert(10)
	list.Insert(15)
	list.Insert(20)
	list.Insert(25)
	list.Insert(30)
	list.Insert(35)
	list.Insert(40)
	list.Insert(45)
	list.Insert(50)

	list.Remove(5)
	list.Remove(10)
	list.Remove(15)
	list.Remove(20)
	list.Remove(25)
	list.Remove(30)
	list.Remove(35)
	list.Remove(40)
	list.Remove(45)
	list.Remove(50)
	if e, a := uint(0), list.Length(); e != a {
		t.Fatalf("was expecting a list of length %v, got %v", e, a)
	}
}
