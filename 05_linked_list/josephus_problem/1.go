package josephus_problem

import (
	"algorithms/05_linked_list/ring"
	"fmt"
)

func solve(total, step int) float32 {
	ringX := cerateRing(total)
	ringX.Step(1)
	for ringX.GetCurrent() != ringX.GetCurrent().GetNext() {
		ringX.Step(step - 1)
		el := ringX.Remove()
		ringX.Step(1)
		el.DisplayLink()
	}
	fmt.Println()
	return ringX.GetCurrent().GetValue()
}

func cerateRing(total int) ring.IRing {
	ringX := ring.NewRing()
	for i := 1; i <= total; i++ {
		ringX.Insert(float32(i))
		ringX.Step(1)
	}
	return ringX
}
