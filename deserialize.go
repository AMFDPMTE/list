package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// Deserialize will take a binary list
// format and return a new List that was initialized from it. If there is any
// thing wrong with tthe bianry input an error will be returned indicating such
//
func Deserialize(data []byte) (list List, err error) {
	// convert our []byte to buffer that can be read
	buffer := bytes.NewReader(data)

	list = List{}

	if len(data) == 0 {
		// there was no data to iterate over
		return
	}

	// we must not have empty list, lets establish the first node
	currentNode := &Node{}

	err = binary.Read(buffer, binary.BigEndian, &currentNode.Value)
	if err != nil {
		return
	}

	// Make the root piont to our first Node
	list.Root = currentNode

	// populate the remaining nodes in the list
	for {
		newNode := &Node{}
		err = binary.Read(buffer, binary.BigEndian, &newNode.Value)
		if err == io.EOF {
			// we have reached the end of our buffer
			err = nil
			break
		} else if err != nil {
			// some other error occured
			return
		}

		if newNode.Value <= currentNode.Value {
			err = fmt.Errorf(
				"New node value was %v, which was smaller or equal to the current node value %v",
				newNode.Value,
				currentNode.Value,
			)
			break
		}

		// The current node needs to point to the next node
		currentNode.Next = newNode

		//	We need to make the Current node the new node then repeat
		currentNode = newNode
	}
	return
}
