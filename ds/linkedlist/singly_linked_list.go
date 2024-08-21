package linkedlist

import "fmt"

type SNode struct {
	value int
	next  *SNode
}

type SinglyLinkedList struct {
	head *SNode
}

func (l *SinglyLinkedList) insert(v int) {
	list := &SNode{value: v, next: nil}

	if l.head == nil {
		l.head = list
	} else {
		t := l.head
		for t.next != nil {
			t = t.next
		}
		t.next = list
	}
}

func (l *SinglyLinkedList) deleteHead() {
	l.head = l.head.next
}

func (l *SinglyLinkedList) deleteTail() {
	t := l.head
	for t.next.next != nil {
		t = t.next
	}
	t.next = nil
}

func (l *SinglyLinkedList) search(v int) bool {
	t := l.head
	for t != nil {
		if t.value == v {
			fmt.Printf("Element with value %d found\n", t.value)
			return true
		}
		t = t.next
	}
	return false
}

func (l *SinglyLinkedList) traverse() {
	fmt.Println("Start of the list")
	t := l.head
	for t != nil {
		fmt.Printf("node value %d, Next node Address %v\n", t.value, t.next)
		t = t.next
	}
	fmt.Println("End of the list")

}

func (l *SinglyLinkedList) sort() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		index := current.next
		for index != nil {
			if current.value > index.value {
				tempValue := current.value
				current.value = index.value
				index.value = tempValue
			}
			index = index.next
		}
		current = current.next
	}
}

func SinglyLinkedListDemo() {
	l := &SinglyLinkedList{nil}

	l.insert(1)
	l.insert(2)
	l.insert(3)
	l.traverse()

	l.search(2)

	l.deleteHead()
	l.deleteTail()
	l.traverse()

	l.insert(123)
	l.insert(12)
	l.insert(142)
	l.insert(15)
	l.insert(1)
	l.insert(15)
	l.sort()
	l.traverse()
}
