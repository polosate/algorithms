package single

import (
	"fmt"

	. "algorithms/15_hash_table"
)

const stepConst = 5

type doubleHashTable struct {
	size      int
	hashArray []*DataItem
	nonItem   *DataItem
}

func newDoubleHashTable(size int) doubleHashTable {
	return doubleHashTable{
		size:      size,
		hashArray: make([]*DataItem, size),
		nonItem:   NewDataItem(-1),
	}
}

func (h *doubleHashTable) find(key int64) *DataItem {
	hashVal := h.hashFunc01(key)
	step := h.hashFunc02(key)
	for h.hashArray[hashVal] != nil {
		if h.hashArray[hashVal].GetKey() == key {
			return h.hashArray[hashVal]
		}
		hashVal += step
		hashVal %= h.size
	}
	return nil
}

func (h *doubleHashTable) insert(item *DataItem) {
	hashVal := h.hashFunc01(item.GetKey())
	step := h.hashFunc02(item.GetKey())
	for h.hashArray[hashVal] != nil && h.hashArray[hashVal].GetKey() != -1 {
		hashVal += step
		hashVal %= h.size
	}
	h.hashArray[hashVal] = item
}

func (h *doubleHashTable) delete(key int64) *DataItem {
	hashVal := h.hashFunc01(key)
	step := h.hashFunc02(key)
	for h.hashArray[hashVal] != nil {
		if h.hashArray[hashVal].GetKey() == key {
			temp := h.hashArray[hashVal]
			h.hashArray[hashVal] = h.nonItem
			return temp
		}
		hashVal += step
		hashVal %= h.size
	}
	return nil
}

func (h *doubleHashTable) hashFunc01(key int64) int {
	return int(key) % len(h.hashArray)
}

func (h *doubleHashTable) hashFunc02(key int64) int {
	return stepConst - int(key)%stepConst
}

func (h *doubleHashTable) display() {
	for i := 0; i < h.size; i++ {
		if h.hashArray[i] != nil {
			fmt.Print(h.hashArray[i].GetKey(), " ")
		} else {
			fmt.Print("** ")
		}
	}
	fmt.Println()
}
