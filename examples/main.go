package main

import (
	"fmt"

	"github.com/AMFDPMTE/list"
)

func main() {
	l1 := list.New(40, 50, 15, 33, 2)

	out := l1.Serialize()
	fmt.Println(out)

	l2, err := list.Deserialize(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(l2)
}
