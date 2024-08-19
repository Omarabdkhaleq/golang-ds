package queue

type Queue interface {
	Enqueue(int) []int
	Dequeue() []int
	IsEmpty() bool
	IsFull() bool
	Peek() int
	Display()
}
