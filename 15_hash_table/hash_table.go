package _5_hash_table

import "fmt"

type hashTable struct {
	size      int
	hashArray []*dataItem
	nonItem   *dataItem
}

func newHashTable(size int) hashTable {
	return hashTable{
		size:      size,
		hashArray: make([]*dataItem, size),
		nonItem:   newDataItem(-1),
	}
}

func (h *hashTable) find(key int64) *dataItem {
	hashVal := h.hashFunc(key)
	for h.hashArray[hashVal] != nil {
		if h.hashArray[hashVal].getKey() == key {
			return h.hashArray[hashVal]
		}
		hashVal++
		hashVal %= h.size
	}
	return nil
}

func (h *hashTable) insert(item *dataItem) {
	hashVal := h.hashFunc(item.getKey())
	for h.hashArray[hashVal] != nil && h.hashArray[hashVal].getKey() != -1 {
		hashVal++
		hashVal %= h.size
	}
	h.hashArray[hashVal] = item
}

func (h *hashTable) delete(key int64) *dataItem {
	hashVal := h.hashFunc(key)
	for h.hashArray[hashVal] != nil {
		if h.hashArray[hashVal].getKey() == key {
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
			fmt.Print(h.hashArray[i].getKey(), " ")
		} else {
			fmt.Print("** ")
		}
	}
	fmt.Println()
}
