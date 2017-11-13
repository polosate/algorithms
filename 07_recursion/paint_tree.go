package _7_recursion

import "fmt"

func PaintTree(size int) {
	paint(0, size)
}

func paint(left, right int) {
	if left == right {
		return
	}

	mid := (left + right) / 2

	for i := left; i <= mid; i++ {
		fmt.Print("-")
	}
	fmt.Print("X")
	for i := mid + 1; i <= left; i++ {
		fmt.Print("-")
	}
	paint(left, mid)
	paint(mid+1, right)
}

8
0 0 0 0 4 0 0 0
0 0 4 0 4 0 4 0
