package queue

import (
	"container/list"
	"fmt"
)

/* Queue => FIFO */

type (
	simpleQueue struct {
		front int
		rear  int
		size  int
		items []int
	}
)

func NewSimpleQueue(size int) simpleQueue {
	return simpleQueue{
		front: -1,
		rear:  -1,
		size:  size,
		items: make([]int, 0, size),
	}
}

func (q simpleQueue) Enqueue(i int) []int {
	if q.IsFull() {
		panic("queue is full")
	}
	if q.front == -1 {
		q.front = 0
	}
	q.rear++
	q.items = append(q.items, i)
	fmt.Printf("inserted: %d\n", i)
	return q.items
}

func (q simpleQueue) Dequeue() []int {

	if q.IsEmpty() {
		panic("queue is empty")
	}

	fmt.Printf("dequeued: %d\n", q.items[q.front])

	b := make([]int, len(q.items)-1)
	_ = copy(b, q.items[1:])
	_ = copy(q.items, b)

	if q.front >= q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front++
	}

	return q.items
}

func (q simpleQueue) IsEmpty() bool {
	if q.front == -1 {
		return true
	}
	return false
}

func (q simpleQueue) IsFull() bool {
	if q.front == 0 && q.rear == q.size-1 {
		return true
	}
	return false
}

func (q simpleQueue) Peek() int {
	return q.items[q.front]
}

func (q simpleQueue) Display() {
	fmt.Println(q.items)
}

func SimpleQueueFromScratch() {

	q := NewSimpleQueue(5)

	fmt.Println(q.IsEmpty())
	fmt.Println(q.IsFull())

	_ = q.Enqueue(1)
	_ = q.Enqueue(2)
	_ = q.Enqueue(3)
	_ = q.Enqueue(4)
	_ = q.Enqueue(5)
	q.Display()

	fmt.Println(q.IsEmpty())
	fmt.Println(q.IsFull())
	fmt.Println(q.Peek())

	//q = q.Enqueue(5) => panic queue is full

	_ = q.Dequeue()
	q.Display()
	_ = q.Dequeue()
	q.Display()
	_ = q.Dequeue()
	q.Display()
	_ = q.Dequeue()
	q.Display()

}

func SimpleQueueLinkedList() {
	q := list.New()
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	front := q.Front()
	q.Remove(front)
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
