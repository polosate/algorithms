package ring

import "fmt"

type IRing interface {
	Insert(value float32)
	Remove() *link
	Peek() *link
	Step(step int)
	IsEmpty() bool
	DisplayRing()
	GetCurrent() *link
}

type ring struct {
	current *link
}

func NewRing() IRing {
	return &ring{
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

func (r *ring) Peek() *link {
	if r.IsEmpty() {
		return nil
	} else {
		return r.current.next
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

func (r *ring) GetCurrent() *link {
	return r.current
}
