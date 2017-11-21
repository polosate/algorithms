package _1_radix_sort

import "fmt"

type RadixArray struct {
	array []int64
	lists []list
}

func newRadixArray(array []int64) RadixArray {
	arr := RadixArray{array: array}
	for i := 0; i < 10; i++ {
		list := NewList()
		arr.lists = append(arr.lists, list)

	}
	return arr
}

func (a *RadixArray) Sort() {
	max := a.max()
	var significantDigit int64 = 1
	for max/significantDigit >= 1 {
		for _, v := range a.array {
			vv := v / significantDigit
			if vv < 1 {
				a.lists[0].InsertLast(v)
			} else {
				vvv := vv % 10
				a.lists[vvv].InsertLast(v)
			}
		}
		j := 0
		for i := range a.lists {
			for !a.lists[i].isEmpty() {
				a.array[j] = a.lists[i].RemoveFirst()
				j++
			}
		}
		significantDigit *= 10
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

func (a *RadixArray) DisplayList() {
	fmt.Println(a.array)
}
