package heap

import "fmt"

type Heap struct {
	a []int
}

func NewHeap() *Heap {
	return &Heap{a: make([]int, 0)}
}

func (h *Heap) swap(a, b int) {
	var i, j int
	for i = 0; i < len(h.a); i++ {
		if h.a[i] == a {
			break
		}
	}
	for j = 0; j < len(h.a); j++ {
		if h.a[j] == b {
			break
		}
	}
	temp := h.a[i]
	h.a[i] = b
	h.a[j] = temp
}

func (h *Heap) heapify(i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < len(h.a) && (h.a)[l] > (h.a)[largest] {
		largest = l
	}
	if r < len(h.a) && (h.a)[r] > (h.a)[largest] {
		largest = r
	}

	if largest != i {
		h.swap(h.a[i], h.a[largest])
		h.heapify(largest)
	}
}

func (h *Heap) insert(value int) {
	h.a = append(h.a, value)
	current := len(h.a) - 1
	for current != 0 {
		parent := (current - 1) / 2
		if h.a[current] > h.a[parent] {
			h.swap(h.a[parent], h.a[current])
			current = parent
		} else {
			break
		}
	}
}

func (h *Heap) deleteHeap(value int) {
	var i int
	for i = 0; i < len(h.a); i++ {
		if h.a[i] == value {
			break
		}
	}

	h.swap(h.a[len(h.a)-1], h.a[i])
	h.a = h.a[:len(h.a)-1]

	if i < len(h.a) {
		h.heapify(i)
	}
}

func (h *Heap) display() {
	fmt.Printf("[")
	for v := range h.a {
		fmt.Printf("%d ", h.a[v])
	}
	fmt.Printf("]")
}

func HeapDemo() {
	h := NewHeap()
	h.insert(1)
	h.insert(23)
	h.insert(53)
	h.insert(12)
	h.insert(3)
	h.insert(6)

	h.display()

	h.deleteHeap(1)

	h.display()
}
