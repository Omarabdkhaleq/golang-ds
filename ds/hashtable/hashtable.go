package hashtable

import (
	"ds/linkedlist"
	"fmt"
	"math"
)

type (
	HashTable struct {
		data []*linkedlist.SinglyLinkedList
	}

	listData struct {
		key string
		val interface{}
	}
)

func new(initSize int) *HashTable {
	return &HashTable{
		make([]*linkedlist.SinglyLinkedList, initSize),
	}
}

func (ht *HashTable) hash(s string) int {
	h := 0
	// hashFunc = char[0] * 31^n-1 + char[1] * 31^n-2 + ... + char[n-1]
	for pos, char := range s {
		h += int(char) * int(math.Pow(31, float64(len(s)-pos+1)))
	}
	return h
}

func (ht *HashTable) index(hash int) int {
	return hash % len(ht.data)
}

// Solved collision by chaining
func (ht *HashTable) insert(k string, v interface{}) {
	index := ht.index(ht.hash(k))

	if ht.data[index] == nil {
		ht.data[index] = &linkedlist.SinglyLinkedList{}
		ht.data[index].Insert(listData{k, v})
		fmt.Printf("Inserted list data %v at new index: %v\n", ht.data[index].Head.Data, index)
	} else {
		node := &ht.data[index].Head
		for {
			if *node != nil {
				d := (*node).Data.(listData)
				if d.key == k {
					d.val = v
					(*node).Data = d
					fmt.Printf("Updated existing value with key %v at shared index: %v\n", k, index)
					break
				}
				node = &(*node).Next
			} else {
				ht.data[index].Insert(listData{k, v})
				fmt.Printf("Inserted new value with key %v at shared index: %v\n", k, index)
				break
			}
		}
	}
}

func (ht *HashTable) get(k string) (result interface{}, ok bool) {
	index := ht.index(ht.hash(k))
	indexList := ht.data[index]

	if indexList == nil {
		fmt.Println("Index not found")
		return nil, false
	}

	node := indexList.Head
	for {
		if node != nil {
			d := node.Data.(listData)
			if d.key == k {
				return d.val, true
			}
		} else {
			return nil, false
		}
		node = node.Next
	}
}

func (ht *HashTable) display() {
	for _, l := range ht.data {
		node := l.Head
		for node != nil {
			fmt.Printf("%v ", node.Data)
			node = node.Next
		}
		fmt.Printf("\n")
	}
}

func Demo() {
	ht := new(5)

	ht.insert("a", 1)
	ht.insert("b", 2)
	ht.insert("c", 3)
	ht.insert("d", 4)
	ht.insert("e", 5)
	ht.insert("f", 6)
	ht.insert("g", 7)
	ht.insert("h", 8)
	ht.insert("i", 9)
	ht.insert("j", 10)
	ht.insert("j", 100)
	ht.display()
	v1, ok1 := ht.get("a")
	fmt.Println(v1, ok1)
	vx, okx := ht.get("w")
	fmt.Println(vx, okx)
}
