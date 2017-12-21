package single

import (
	"fmt"

	. "algorithms/15_hash_table"
)

type hashTable struct {
	size      int
	curSize   int
	hashArray []*DataItem
	nonItem   *DataItem
}

func newHashTable(size int) hashTable {
	return hashTable{
		size:      size,
		curSize:   0,
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
	if float32(h.curSize)/float32(h.size) > 0.5 {
		h.rehash()
	}
	hashVal := h.hashFunc(item.GetKey())
	for h.hashArray[hashVal] != nil && h.hashArray[hashVal].GetKey() != -1 {
		hashVal++
		hashVal %= h.size
	}
	h.hashArray[hashVal] = item
	h.curSize++
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
	return int(key) % h.size
}

func (h *hashTable) rehash() {
	newSize := getPrime(h.size * 2)
	newHashTable := newHashTable(newSize)
	for i := 0; i < h.size; i++ {
		if h.hashArray[i] != nil && h.hashArray[i].GetKey() != -1 {
			newHashTable.insert(h.hashArray[i])
		}
		if h.hashArray[i].GetKey() == -1 {

		}
	}
	h.hashArray = newHashTable.hashArray
	h.size = newSize
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

func getPrime(min int) int {
	for i := min + 1; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

func isPrime(n int) bool {
	for j := 2; j*j < n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}
