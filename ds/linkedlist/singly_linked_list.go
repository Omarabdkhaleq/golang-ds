package linkedlist

import "fmt"

type SNode struct {
	Data interface{}
	Next *SNode
}

type SinglyLinkedList struct {
	Head *SNode
}

func (l *SinglyLinkedList) Insert(v interface{}) {
	list := &SNode{Data: v, Next: nil}

	if l.Head == nil {
		l.Head = list
	} else {
		t := l.Head
		for t.Next != nil {
			t = t.Next
		}
		t.Next = list
	}
}

func (l *SinglyLinkedList) deleteHead() {
	l.Head = l.Head.Next
}

func (l *SinglyLinkedList) deleteTail() {
	t := l.Head
	for t.Next.Next != nil {
		t = t.Next
	}
	t.Next = nil
}

func (l *SinglyLinkedList) search(v int) bool {
	t := l.Head
	for t != nil {
		if t.Data == v {
			fmt.Printf("Element with Data %d found\n", t.Data)
			return true
		}
		t = t.Next
	}
	return false
}

func (l *SinglyLinkedList) traverse() {
	fmt.Println("Start of the list")
	t := l.Head
	for t != nil {
		fmt.Printf("node Data %d, Next node Address %v\n", t.Data, t.Next)
		t = t.Next
	}
	fmt.Println("End of the list")

}

// TODO: find solution to sort interface{}
//func (l *SinglyLinkedList) sort() {
//	if l.head == nil {
//		return
//	}
//	current := l.head
//	for current != nil {
//		index := current.Next
//		for index != nil {
//			if current.Data > index.Data {
//				tempValue := current.Data
//				current.Data = index.Data
//				index.Data = tempValue
//			}
//			index = index.Next
//		}
//		current = current.Next
//	}
//}

func SinglyLinkedListDemo() {
	l := &SinglyLinkedList{nil}

	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.traverse()

	l.search(2)

	l.deleteHead()
	l.deleteTail()
	l.traverse()

	l.Insert(123)
	l.Insert(12)
	l.Insert(142)
	l.Insert(15)
	l.Insert(1)
	l.Insert(15)
	//l.sort()
	l.traverse()
}
