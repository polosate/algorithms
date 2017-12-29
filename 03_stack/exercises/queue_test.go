package exercises

import "testing"

func TestQueue(t *testing.T) {

	q := newQueue(5)
	q.add(5)
	q.add(10)
	if data, _ := q.remove(); data != 5 {
		t.Error()
	}
	q.add(7)

	if data, _ := q.remove(); data != 10 {
		t.Error()
	}
	if data, _ := q.remove(); data != 7 {
		t.Error()
	}
}
