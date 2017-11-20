package _1_radix_sort

type RadixArray struct {
	array []int64
	list0 list
	list1 list
	list2 list
	list3 list
	list4 list
	list5 list
	list6 list
	list7 list
	list8 list
	list9 list
}

func newRadixArray(array []int64) RadixArray {
	return RadixArray{array: array}
}

// 435 123 045 67 911 234
func (a *RadixArray) Sort() {
	for _, v := range a.array {

	}
}

func (a *RadixArray) max() int64 {
	max := a.array[0]
	for i := 0; i < len(a.array); i++ {
		if a.array[i] > max {
			max = a.array[i]
		}
	}
	return max
}
