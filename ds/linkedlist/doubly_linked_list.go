package linkedlist

import "fmt"

type DNode struct {
	value interface{}
	next  *DNode
	prev  *DNode
}

type DoublyLinkedList struct {
	head *DNode
}

func (d *DoublyLinkedList) insertFront(v int) {
	n := &DNode{value: v, next: nil, prev: nil}

	if d.head == nil {
		d.head = n
	} else {
		n.next = d.head
		d.head.prev = n
		d.head = n
	}
}

func (d *DoublyLinkedList) insertBetween(firstNode, secondNode *DNode, v int) {
	if d.head == nil {
		panic("List is empty")
	}
	n := &DNode{value: v, next: nil, prev: nil}
	firstNode.next = n
	n.prev = firstNode
	secondNode.prev = n
	n.next = secondNode
}

func (d *DoublyLinkedList) insertTail(v int) {
	n := &DNode{value: v, next: nil, prev: nil}

	if d.head == nil {
		d.head = n
	} else {
		node := d.head
		for node.next != nil {
			node = node.next
		}
		node.next = n
		n.prev = node

	}
}

func (d *DoublyLinkedList) deleteFront() {
	if d.head == nil {
		panic("List is empty")
	}
	d.head.next.prev = nil
	d.head = d.head.next
}

func (d *DoublyLinkedList) deleteTail() {
	node := d.head
	for node.next != nil {
		node = node.next
	}
	node.prev.next = nil
	node = nil
}

func (d *DoublyLinkedList) deleteBetween(firstNode, secondNode *DNode) {
	if d.head == nil {
		panic("List is empty")
	}
	firstNode.next = secondNode
	secondNode.prev = firstNode
}

func (d *DoublyLinkedList) search(v int) bool {
	//TODO implement me
	panic("implement me")
}

func (d *DoublyLinkedList) traverse() {
	fmt.Println("Start of the list")
	t := d.head
	for t != nil {
		fmt.Printf("node Data %d,Next node Address %v, Prev node address %v\n", t.value, t.next, t.prev)
		t = t.next
	}
	fmt.Println("End of the list")
}

func (d *DoublyLinkedList) sort() {
	//TODO implement me
	panic("implement me")
}

func DoublyLinkedListDemo() {
	d := &DoublyLinkedList{nil}
	d.insertFront(1)
	d.insertFront(2)
	d.insertFront(3)
	d.insertBetween(d.head.next, d.head.next.next, 10)
	d.insertTail(111)
	d.insertFront(99)
	d.deleteFront()
	d.insertTail(222)
	d.deleteTail()
	d.insertBetween(d.head.next, d.head.next.next, 10000)
	d.deleteBetween(d.head.next, d.head.next.next.next)
	d.traverse()

}
