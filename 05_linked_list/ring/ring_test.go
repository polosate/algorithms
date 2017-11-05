package ring

import "testing"

func TestRingInsert(t *testing.T) {
	ring := NewRing()
	ring.Insert(1)
	ring.Insert(2)
	ring.Insert(3)
	ring.Insert(4)
	ring.Insert(5)
	//ring.DisplayRing()
	el := ring.Peek()
	if el.GetValue() != 5 {
		t.Error("Expected el", 5, "have", el.GetValue())
	}
}

func TestRingStep(t *testing.T) {
	ring := NewRing()
	ring.Insert(1)
	ring.Insert(2)
	ring.Insert(3)
	ring.Insert(4)
	ring.Insert(5)
	ring.Insert(6)

	if ring.GetCurrent().GetValue() != 1 {
		t.Error("Expected current", 1, "have", ring.GetCurrent().GetValue())
	}

	ring.Step(3)
	if ring.GetCurrent().GetValue() != 4 {
		t.Error("Expected current", 4, "have", ring.GetCurrent().GetValue())
	}
	//ring.DisplayRing()
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
	if ring.GetCurrent() != nil {
		t.Error("Expected current", nil, "have", ring.GetCurrent())
	}

	ring.Insert(2)
	ring.Insert(3)
	ring.Insert(4)
	ring.Insert(5)

	el = ring.Remove()
	if el.GetValue() != 5 {
		t.Error("Expected el", 5, "have", el.GetValue())
	}
	if ring.GetCurrent().GetNext().GetValue() != 4 {
		t.Error("Expected current", 4, "have", ring.GetCurrent().GetNext().GetValue())
	}
	//ring.DisplayRing()

}
