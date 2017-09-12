package main

import "fmt"

func main() {
	list := List{}
	list.Insert(2)
	list.Insert(40)
	list.Insert(50)
	list.Insert(15)
	list.Insert(33)

	data := list.Serialize()
	fmt.Println(data)
	Deserialize(data)
	// list.Remove(2)
	// list.Remove](10)
	// list.Remove(25)
	// list.Remove(40)
	fmt.Println(list)
}
