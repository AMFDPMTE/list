package main

import (
	"bytes"
	"encoding/binary"
)

// Deserialize will take a binary list
// format and return a new List that was initialized from it. If there is any
// thing wrong with tthe bianry input an error will be returned indicating such
//
func Deserialize(data []byte) (list List, err error) {
	var value uint16
	buffer := bytes.NewReader(data)

	err = binary.Read(buffer, binary.BigEndian, &value)

	if err != nil {
		return
	}

	// TODO: this is only got us the first value; it also didn't maek us a list
	return
}
