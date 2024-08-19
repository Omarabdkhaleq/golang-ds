package queue

import (
	"container/list"
	"fmt"
)

type (
	circularQueue struct {
		front int
		rear  int
		size  int
		items *list.List
	}
)

func NewCircularQueue(size int) circularQueue {
	return circularQueue{
		-1,
		-1,
		size,
		list.New(),
	}
}

func (c *circularQueue) Enqueue(i int) {

	if c.IsFull() {
		panic("queue is full")
	}

	if c.front == -1 {
		c.front = 0
	}

	if (c.front != 0 && c.rear == c.size-1) || (c.rear < c.front && c.rear != -1) {
		c.rear = 0
		c.items.PushFront(i)
	} else {
		c.items.PushBack(i)
	}

	c.rear++

	fmt.Printf("Enqueued: %d, at rear %d, front: %d\n", i, c.rear, c.front)

}

func (c *circularQueue) Dequeue() {
	if c.IsEmpty() {
		panic("queue is empty")
	}
	elem := c.items.Front()
	c.items.Remove(elem)

	fmt.Printf("Dequeued: %d, at rear %d, front: %d\n", elem.Value, c.rear, c.front)

	if c.front == c.rear {
		c.front = -1
		c.rear = -1
	} else {
		c.front++
	}
}

func (c *circularQueue) IsEmpty() bool {
	if c.front == -1 && c.rear == -1 {
		return true
	}
	return false
}

func (c *circularQueue) IsFull() bool {
	if c.front == 0 && c.rear == c.size-1 {
		return true
	}
	if c.front == c.rear+1 {
		return true
	}
	return false
}

func (c *circularQueue) Peek() *list.Element {
	return c.items.Front()
}

func (c *circularQueue) Display() {
	q := make([]int, 0, c.size)
	for e := c.items.Front(); e != nil; e = e.Next() {
		q = append(q, e.Value.(int))
	}
	fmt.Printf("Circular Queue: %v\n", q)
}

func CircularQueueFromScratch() {
	c := NewCircularQueue(5)

	c.Enqueue(1)
	c.Enqueue(2)
	c.Enqueue(3)
	c.Enqueue(4)
	c.Enqueue(5)
	c.Display()

	c.Dequeue()
	c.Dequeue()
	c.Dequeue()
	c.Dequeue()
	c.Display()

	c.Enqueue(99)
	c.Enqueue(88)
	c.Enqueue(77)
	c.Dequeue()
	c.Enqueue(66)
	c.Display()

}
