package main

import "fmt"

// List is a reference to the start of a list (aka the start of a bunch of nodes)
type List struct {
	Root *Node
}

// Insert will insert a new uint value into the list at the appropriate location
// If the value was already in the list (duplicate) an error will be returned.
func (l *List) Insert(n uint) error {
	newNode := Node{Value: n}

	if l.Root == nil {
		// Establish the first node in the list
		l.Root = &newNode
		return nil // Node inserted we can exit!
	}

	// If we reached here, this means we had at least one node (root node) in the list
	var current, next *Node
	current = l.Root

	if newNode.Value < current.Value {
		// newNode should be the first in the list
		newNode.Next = current
		l.Root = &newNode
		return nil
	}

	for current != nil {
		next = current.Next

		if newNode.Value == current.Value {
			return fmt.Errorf("duplicate value %v cannot be inserted", newNode.Value)
		}

		if next == nil {
			// our new node is the last one in the list
			current.Next = &newNode
			return nil // inserted!
		}

		if newNode.Value > current.Value && newNode.Value < next.Value {
			// our new node goes in between two of the old ones
			newNode.Next = next
			current.Next = &newNode
			return nil // inserted!
		}

		current = next // progress to the next node in the list
	}

	return nil
}

// Length returns the number of Nodes in a given array
func (l List) Length() (length uint) {
	current := l.Root

	for current != nil {
		length++
		current = current.Next
	}
	return
}

// Contain checks whether or not the value is in a given list
func (l List) Contain(val uint) (c bool) {
	current := l.Root

	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}
	return
}

// Remove the Node from the list
func (l *List) Remove(n uint) error {
	return nil
}

// String will return a comma seperated representation of the list
func (l List) String() string {
	s := "<start>"
	if l.Root != nil {
		n := l.Root
		for n != nil {
			if n.Next == nil {
				// last one
				s = s + fmt.Sprintf("%v<end>", n.Value)
			} else {
				// not last
				s = s + fmt.Sprintf("%v, ", n.Value)
			}
			n = n.Next
		}
	}
	return s
}
