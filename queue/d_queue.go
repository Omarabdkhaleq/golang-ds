package queue

import (
	"fmt"
)

type DQueue struct {
	size  int
	front int
	rear  int
	items []int
}

func newDQueue(size int) *DQueue {
	return &DQueue{
		size:  size,
		front: -1,
		rear:  -1,
		items: make([]int, size),
	}
}

func (q *DQueue) frontEnqueue(v int) {
	if q.isFull() {
		panic("queue is full")
	}

	if q.front < 1 {
		q.front = q.size - 1
	} else {
		q.front--
	}
	q.items[q.front] = v
}

func (q *DQueue) rearEnqueue(v int) {
	if q.isFull() {
		panic("queue is full")
	} else if q.rear == q.size-1 {
		q.rear = 0
	} else {
		q.rear++
	}

	q.items[q.rear] = v
}

func (q *DQueue) frontDequeue() {
	if q.isEmpty() {
		panic("queue is empty")
	}

	q.items[q.front] = 0

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else if q.front == q.size-1 {
		q.front = 0
	} else {
		q.front++
	}
}

func (q *DQueue) rearDequeue() {
	if q.isEmpty() {
		panic("queue is empty")
	}

	q.items[q.rear] = 0

	if q.rear == q.front {
		q.front = -1
		q.rear = -1
	} else if q.rear == 0 {
		q.rear = q.size - 1
	} else {
		q.rear--
	}
}

func (q *DQueue) isEmpty() bool {
	return q.front == -1 || (q.front == 0 && q.rear == q.size-1)
}

func (q *DQueue) isFull() bool {
	if q.front == q.rear+1 {
		return true
	}
	return false
}

func (q *DQueue) display() {
	fmt.Printf("DQueue: %v\n", q.items)
}

func DQueueFromScratch() {
	q := newDQueue(5)

	q.frontEnqueue(1)
	q.frontEnqueue(2)
	q.rearEnqueue(5)
	q.rearEnqueue(4)
	q.display()

	q.frontDequeue()
	q.rearDequeue()
	q.rearDequeue()
	q.display()
}
