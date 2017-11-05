package ring

import "testing"

func TestRingInsert(t *testing.T) {
	ring := NewRing()
	ring.Insert(1)
	ring.Insert(2)
	ring.Insert(3)
	ring.Insert(4)
	ring.Insert(5)
	ring.DisplayRing()
}

func TestRingStep(t *testing.T) {
	ring := NewRing()
	ring.Insert(1)
	ring.Insert(2)
	ring.Insert(3)
	ring.Insert(4)
	ring.Insert(5)
	ring.Insert(6)

	if ring.current.GetValue() != 6 {
		t.Error("Expected current", 6, "have", ring.current.GetValue())
	}

	ring.Step(3)
	if ring.current.GetValue() != 3 {
		t.Error("Expected current", 3, "have", ring.current.GetValue())
	}
	ring.DisplayRing()
}

func TestRingRemove(t *testing.T) {
	ring := NewRing()
	el := ring.Remove()
	if el != nil {
		t.Error("Expected el", nil, "have", el)
	}
	ring.Insert(1)
	el = ring.Remove()
	if el.GetValue() != 1 {
		t.Error("Expected el", 1, "have", el.GetValue())
	}
	if ring.current != nil {
		t.Error("Expected current", nil, "have", ring.current)
	}

	ring.Insert(2)
	ring.Insert(3)
	ring.Insert(4)
	ring.Insert(5)

	el = ring.Remove()
	if el.GetValue() != 2 {
		t.Error("Expected el", 2, "have", el.GetValue())
	}
	if ring.current.next.GetValue() != 3 {
		t.Error("Expected current", 3, "have", ring.current.next.GetValue())
	}
	ring.DisplayRing()

}
