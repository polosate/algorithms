package _5_hash_table

import (
	"testing"
)

func TestInsert(t *testing.T) {
	h := newHashTable(30)
	h.insert(newDataItem(31))
	h.insert(newDataItem(32))
	h.insert(newDataItem(991))

	if h.find(31).getKey() != 31 {
		t.Error("want", 31, "have", h.find(31).getKey())
	}
	if h.find(100) != nil {
		t.Error("want", nil, "have", h.find(100))
	}

	h.display()

	if h.delete(31).getKey() != 31 {
		t.Error("want", 31, "have", h.find(31).getKey())
	}
	if h.delete(100) != nil {
		t.Error("want", nil, "have", h.find(100))
	}

	h.display()
	h.insert(newDataItem(61))
	h.display()
}
