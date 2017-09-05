package main

import (
	"bytes"
	"encoding/binary"
)

// Serialize will serialize the list in our binary format
func (l List) Serialize() (data []byte) {
	if l.Root == nil {
		// We had a list without anything in it; we just return empty []byte
		return
	}

	//We have something in the list
	current := l.Root
	for {
		// We need buffer to write the current uint16 value into
		values := new(bytes.Buffer)
		err := binary.Write(values, binary.BigEndian, current.Value)
		if err != nil {
			panic(err)
		}
		// The buffer will now have 2 bytes; we append those to our data variable
		// that we will return
		data = append(data, values.Bytes()...)
		// CHeck if next is pointing to nil
		if current.Next == nil {
			break
		}
		//We have another node to continue; move current to the next node
		current = current.Next
	}
	return
}
