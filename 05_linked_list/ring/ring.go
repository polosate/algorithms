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

func (r *ring) Remove() *link {
	if r.IsEmpty() {
		return nil
	} else if r.current == r.current.next {
		temp := r.current
		r.current = nil
		return temp
	} else {
		temp := r.current.next
		r.current.next = temp.next
		return temp
	}
}

func (r *ring) Step(step int) {
	for i := 0; i < step; i++ {
		r.current = r.current.next
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
		current := r.current
		current.DisplayLink()
		current = current.next
		for current != r.current {
			current.DisplayLink()
			current = current.next
		}
	}
	fmt.Println("\n--------")
}
