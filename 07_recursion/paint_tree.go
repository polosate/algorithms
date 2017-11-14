package _7_recursion

import (
	"fmt"
	"math"
)

func PaintTree(deep float64) {
	res := make([][]string, int(deep)+1)
	makeTree(0, int(math.Pow(2, deep))-1, 0, res)
	display(res)
}

func makeTree(left, right, deep int, res [][]string) {
	if left == right {
		res[deep] = append(res[deep], "X")
		return
	}
	mid := (left + right) / 2
	for i := left; i <= mid; i++ {
		res[deep] = append(res[deep], "-")
	}
	res[deep] = append(res[deep], "X")
	for i := mid + 2; i <= right; i++ {
		res[deep] = append(res[deep], "-")
	}
	makeTree(left, mid, deep+1, res)
	makeTree(mid+1, right, deep+1, res)
}

func display(tree [][]string) {
	for _, v := range tree {
		for _, k := range v {
			fmt.Print(k)
		}
		fmt.Println()
	}
}
