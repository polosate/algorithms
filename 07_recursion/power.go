package _7_recursion

func Power(x, n int) int {
	if n == 1 {
		return x
	} else if n%2 == 0 {
		return Power(x*x, n/2)
	} else {
		return x * Power(x*x, n/2)
	}
}
