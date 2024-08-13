package main

import (
	"ds/ds"
	"fmt"
)

func main() {

	s := ds.NewStack(4)
	s = s.Push(1)
	fmt.Println(s)
	s = s.Push(2)
	fmt.Println(s)
	fmt.Println(s.Peek())
	s = s.Pop()
	fmt.Println(s)
	s = s.Pop()
	fmt.Println(s)

}
