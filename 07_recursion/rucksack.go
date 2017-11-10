package _7_recursion

type rucksack struct {
	size  int
	sum   int
	array []int
}

func NewRucksack(sum int, array []int) rucksack {
	return rucksack{
		size:  len(array),
		sum:   sum,
		array: array,
	}
}

func (r *rucksack) Do() []int {
	res := make([]int, 0, len(r.array))
	return r.do(r.sum, r.array, res)
}

func (r *rucksack) do(sum int, array []int, res []int) []int {
	if sum-array[0] == 0 {
		res = append(res, array[0])
		return res
	} else {
		if sum-array[0] > 0 {
			sum = sum - array[0]
			res = append(res, array[0])
			array = array[1:]
			return r.do(sum, array, res)
		} else {
			if len(array) > 1 {
				return r.do(sum, array[1:], res)
			} else {
				if len(r.array) > 1 {
					r.array = r.array[1:]
					newRes := make([]int, 0, r.size)
					newArray := r.array
					return r.do(r.sum, newArray, newRes)
				} else {
					return []int{0, 0, 0}
				}
			}
		}
	}
}
