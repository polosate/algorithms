package _7_recursion

func triangle(n int) int {
	sum := 0
	for n > 0 {
		sum += n
		n--
	}
	return sum
}

func recursionTriangle(n int) int {
	if n == 1 {
		return 1
	}
	return recursionTriangle(n-1) + n
}
