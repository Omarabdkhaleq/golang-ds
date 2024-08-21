package linkedlist

type DNode struct {
	value int
	next  *DNode
	prev  *DNode
}

type DoublyLinkedList struct {
	head *DNode
}

func (d DoublyLinkedList) insert(v int) {
	n := &DNode{value: v, next: nil}
	if d.head == nil {
		d.head = n
	}
	tmp := d.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = n
}

func (d DoublyLinkedList) deleteHead() {
	//TODO implement me
	panic("implement me")
}

func (d DoublyLinkedList) deleteTail() {
	//TODO implement me
	panic("implement me")
}

func (d DoublyLinkedList) search(v int) bool {
	//TODO implement me
	panic("implement me")
}

func (d DoublyLinkedList) traverse() {
	//TODO implement me
	panic("implement me")
}

func (d DoublyLinkedList) sort() {
	//TODO implement me
	panic("implement me")
}
