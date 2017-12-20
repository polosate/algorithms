package hashChain

import (
	. "algorithms/15_hash_table"
	"fmt"
)

type hashChain struct {
	size      int
	hashArray []list
}

func newHashChain(size int) hashChain {
	var lists []list
	for i := 0; i < size; i++ {
		lists = append(lists, newList())
	}
	return hashChain{
		size:      size,
		hashArray: lists,
	}
}

func (h *hashChain) find(key int64) *DataItem {
	hashVal := h.hashFunc(key)
	return h.hashArray[hashVal].find(key)
}

func (h *hashChain) insert(item *DataItem) {
	hashVal := h.hashFunc(item.GetKey())
	h.hashArray[hashVal].insertFirst(item)
}

func (h *hashChain) delete(key int64) *DataItem {
	hashVal := h.hashFunc(key)
	return h.hashArray[hashVal].delete(key)
}

func (h *hashChain) hashFunc(key int64) int {
	return int(key) % len(h.hashArray)
}

func (h *hashChain) display() {
	for i := 0; i < h.size; i++ {
		h.hashArray[i].display()
	}
	fmt.Println("=====end=======")
}
