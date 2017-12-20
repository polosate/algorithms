package hashChain

import (
	"testing"

	. "algorithms/15_hash_table"
)

func TestHashChain(t *testing.T) {
	h := newHashChain(10)
	h.insert(NewDataItem(91))
	h.insert(NewDataItem(30))
	h.insert(NewDataItem(40))
	h.display()
	h.delete(40)
	h.delete(30)
	h.display()
}
