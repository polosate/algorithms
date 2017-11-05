package ring

import "fmt"

type ring struct {
	current *link
}

func NewRing() ring {
	return ring{
		current: nil,
	}
}

func (r *ring) Insert(value float32) {
	link := NewLink(value)
	if r.current == nil {
		r.current = link
		r.current.next = link
	} else {
		link.next = r.current.next
		r.current.next = link
		r.current = link
	}
}

func (r *ring) IsEmpty() bool {
	return r.current == nil
}

func (r *ring) DisplayRing() {
	if r.IsEmpty() {
		fmt.Println("ring is empty")
	} else if r.current == r.current.next {
		r.current.DisplayLink()
	} else {
		current := r.current.next
		current.DisplayLink()
		current = current.next
		for current != r.current.next {
			current.DisplayLink()
			current = current.next
		}
	}
	fmt.Println("\n--------")
}
