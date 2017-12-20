package single

import (
	"testing"

	. "algorithms/15_hash_table"
)

func TestInsert(t *testing.T) {
	h := newHashTable(30)
	h.insert(NewDataItem(31))
	h.insert(NewDataItem(32))
	h.insert(NewDataItem(991))

	if h.find(31).GetKey() != 31 {
		t.Error("want", 31, "have", h.find(31).GetKey())
	}
	if h.find(100) != nil {
		t.Error("want", nil, "have", h.find(100))
	}

	h.display()

	if h.delete(31).GetKey() != 31 {
		t.Error("want", 31, "have", h.find(31).GetKey())
	}
	if h.delete(100) != nil {
		t.Error("want", nil, "have", h.find(100))
	}

	h.display()
	h.insert(NewDataItem(61))
	h.display()
}
