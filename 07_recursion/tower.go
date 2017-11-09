package _7_recursion

import (
	"fmt"

	stack "algorithms/03_stack"
)

type tower struct {
	count int
	size  int
	from  stack.IntStack
	inter stack.IntStack
	to    stack.IntStack
}

func NewTower(size int) tower {
	tower := tower{
		count: 0,
		size:  size,
		from:  stack.NewIntStack(size),
		inter: stack.NewIntStack(size),
		to:    stack.NewIntStack(size),
	}
	for i := size; i > 0; i-- {
		tower.from.Push(int64(i))
	}
	return tower
}

func (t *tower) Move() {
	t.move(t.size, &t.from, &t.inter, &t.to)
}

func (t *tower) move(top int, from, inter, to *stack.IntStack) {
	if top == 1 {
		el, _ := from.Pop()
		to.Push(el)
		t.count++
	} else {
		t.move(top-1, from, to, inter)
		t.move(1, from, inter, to)
		t.move(top-1, inter, from, to)
	}
}

func (t *tower) DisplayTowers() {
	fmt.Print("A: ")
	t.from.DisplayStack()
	fmt.Print("B: ")
	t.inter.DisplayStack()
	fmt.Print("C: ")
	t.to.DisplayStack()
}
