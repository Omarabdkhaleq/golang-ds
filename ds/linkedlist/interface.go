package linkedlist

type (
	LinkedList interface {
		insert(v int)
		deleteHead()
		deleteTail()
		search(v int) bool
		traverse()
		sort()
	}
)
