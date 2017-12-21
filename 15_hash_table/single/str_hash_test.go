package single

import (
	"testing"

	. "algorithms/15_hash_table"
)

func TestHashString(t *testing.T) {
	h := newStrHash(10)
	h.insert(NewStrDataItem("test"))
	h.insert(NewStrDataItem("cat"))

	w := h.find("bad")
	if w != nil {
		t.Error()
	}

	w = h.find("cat")
	if w.GetKey() != "cat" {
		t.Error()
	}

	d := h.delete("cat")
	if d.GetKey() != "cat" {
		t.Error()
	}

	w = h.find("cat")
	if w != nil {
		t.Error()
	}
}
