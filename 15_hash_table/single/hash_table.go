package single

import (
	"fmt"

	. "algorithms/15_hash_table"
)

type hashTable struct {
	size      int
	hashArray []*DataItem
	nonItem   *DataItem
}

func newHashTable(size int) hashTable {
	return hashTable{
		size:      size,
		hashArray: make([]*DataItem, size),
		nonItem:   NewDataItem(-1),
	}
}

func (h *hashTable) find(key int64) *DataItem {
	hashVal := h.hashFunc(key)
	for h.hashArray[hashVal] != nil {
		if h.hashArray[hashVal].GetKey() == key {
			return h.hashArray[hashVal]
		}
		hashVal++
		hashVal %= h.size
	}
	return nil
}

func (h *hashTable) insert(item *DataItem) {
	hashVal := h.hashFunc(item.GetKey())
	for h.hashArray[hashVal] != nil && h.hashArray[hashVal].GetKey() != -1 {
		hashVal++
		hashVal %= h.size
	}
	h.hashArray[hashVal] = item
}

func (h *hashTable) delete(key int64) *DataItem {
	hashVal := h.hashFunc(key)
	for h.hashArray[hashVal] != nil {
		if h.hashArray[hashVal].GetKey() == key {
			temp := h.hashArray[hashVal]
			h.hashArray[hashVal] = h.nonItem
			return temp
		}
		hashVal++
		hashVal %= h.size
	}
	return nil
}

func (h *hashTable) hashFunc(key int64) int {
	return int(key) % len(h.hashArray)
}

func (h *hashTable) display() {
	for i := 0; i < h.size; i++ {
		if h.hashArray[i] != nil {
			fmt.Print(h.hashArray[i].GetKey(), " ")
		} else {
			fmt.Print("** ")
		}
	}
	fmt.Println()
}
