package single

import (
	"fmt"

	. "algorithms/15_hash_table"
)

type strHash struct {
	size      int
	hashArray []*StrDataItem
	nonItem   *StrDataItem
}

func newStrHash(size int) strHash {
	return strHash{
		size:      size,
		hashArray: make([]*StrDataItem, size),
		nonItem:   NewStrDataItem(""),
	}
}

func (h *strHash) find(key string) *StrDataItem {
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

func (h *strHash) insert(item *StrDataItem) {
	hashVal := h.hashFunc(item.GetKey())
	for h.hashArray[hashVal] != nil && h.hashArray[hashVal].GetKey() != "" {
		hashVal++
		hashVal %= h.size
	}
	h.hashArray[hashVal] = item
}

func (h *strHash) delete(key string) *StrDataItem {
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

func (h *strHash) hashFunc(key string) int {
	hashVal := 0
	for _, v := range key {
		hashVal = (hashVal*26 + int(v)) % h.size
	}

	return hashVal
}

func (h *strHash) display() {
	for i := 0; i < h.size; i++ {
		if h.hashArray[i] != nil {
			fmt.Print(h.hashArray[i].GetKey(), " ")
		} else {
			fmt.Print("** ")
		}
	}
	fmt.Println()
}
