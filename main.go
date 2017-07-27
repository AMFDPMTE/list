package main

import "fmt"

func main() {
	list := List{}
	list.Insert(10)
	list.Insert(5)
	list.Insert(20)
	list.Insert(30)
	list.Insert(25)
	fmt.Println(list)
}
