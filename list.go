package main

import "fmt"

// List is a reference to the start of a list (aka the start of a bunch of nodes)
type List struct {
	Root *Node
}

// New is a variadic function meaning it can take any number of argumnents as
// input and create a new ordered list with them. The values probided may be out
// of order (unordered). The list we produce will be ordered.
func New(values ...uint16) (l *List) {
	l = &List{} // init the ptr
	// TODO: implement
	return
}

// Insert will insert a new uint value into the list at the appropriate location
// If the value was already in the list (duplicate) an error will be returned.
func (l *List) Insert(n uint16) error {
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

// ValueAt returns the value at the specified position index. The list
// is zero indexed meaning the first value is at 0. If an out of range index is
// provided, an error is returned.
func (l List) ValueAt(index int) (value uint16, err error) {
	// TODO: implement
	return
}

// Contain checks whether or not the value is in a given list
func (l List) Contain(val uint16) (c bool) {
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
// why is *List(pointers?)
func (l *List) Remove(n uint16) error {

	if l.Root == nil {
		return fmt.Errorf("value does not exist in List or the list is empty")
	}

	prev, curr := l.Root, l.Root

	if curr.Value == n {
		l.Root = curr.Next
		curr = nil
		return nil
	}

	for {
		// 1. Move cur to Next
		curr = curr.Next

		// 2. if cur is nil we have reached the end of the list and the value has
		//    not been found; thus we error. If cur is not nil continue to step 3.
		if curr == nil {
			return fmt.Errorf("value did not exist in List")
		}
		// 3. if cur equals our n value, we found the value we wanted to remove. We
		//    need to set prev.Next to cur.Next and cur.Next to nil; we have now
		//    removed the value and can return. If cur did not equal our n value we
		//    have not foud the value we wish to remove; proceed to step 4.
		if curr.Value == n {
			prev.Next = curr.Next
			curr.Next = nil
			return nil
		}
		// 4. We need to get prev and current to the same node (this was our initial
		//    conditions for looping) - thus we make prev prev.Next. Prev and Cur
		//    are now pointing at the same thing and we can loop.
		prev = prev.Next

	}
}

// String will return a comma seperated representation of the list
func (l List) String() string {
	s := "<start>"
	e := "<end>"

	if l.Root == nil {
		s = s + e
	}

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
