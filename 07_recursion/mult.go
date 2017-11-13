package _7_recursion

func multiply(x, y int) int {
	if y == 0 {
		return 0
	}
	if y == 1 {
		return x
	}
	return x + multiply(x, y-1)
}
