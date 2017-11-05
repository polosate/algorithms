package ring

import "testing"

func TestRingInsert(t *testing.T) {
	ring := NewRing()
	ring.DisplayRing()
	ring.Insert(1)
	ring.DisplayRing()
	ring.Insert(2)
	ring.DisplayRing()
	ring.Insert(3)
	ring.DisplayRing()
	ring.Insert(4)
	ring.DisplayRing()
	ring.Insert(5)
	ring.DisplayRing()
	ring.Insert(6)
	ring.DisplayRing()
}
