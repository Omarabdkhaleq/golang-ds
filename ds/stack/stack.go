package stack

import "fmt"

/* STACK => LIFO */

type (
	stack struct {
		top   int
		stack []int
	}
	Stack interface {
		Push(int) stack
		Pop() stack
		IsEmpty() bool
		IsFull() bool
		Peek() int
	}
)

func NewStack(size int) Stack {
	return stack{
		top:   -1,
		stack: make([]int, 0, size),
	}
}

func (s stack) Push(i int) stack {
	if s.IsFull() {
		panic("stack is full")
	}
	s.stack = append(s.stack, i)
	s.top++
	return s
}

func (s stack) Pop() stack {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	s.stack = s.stack[:s.top]
	s.top--
	return s
}

func (s stack) IsEmpty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

func (s stack) IsFull() bool {
	if len(s.stack) == 0 {
		return false
	}
	if s.top == len(s.stack) {
		return true
	}
	return false
}

func (s stack) Peek() int {
	return s.stack[s.top]
}

func StackFromScratch() {
	s := NewStack(4)
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
