package _7_recursion

func knuthSeq(limit int) []int {
	var res []int
	k := 0
	for {
		k = 3*k + 1
		if k < limit {
			res = append(res, k)
		} else {
			break
		}
	}
	return res
}
