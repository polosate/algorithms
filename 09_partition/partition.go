package _9_partition

import "fmt"

type PartArray struct {
	array []int64
}

func NewPartArray(array []int64) PartArray {
	return PartArray{
		array: array,
	}
}

func (a *PartArray) Partition(left, right int, pivot int64) int {
	leftPtr := left - 1
	rightPtr := right + 1
	var i int

	for {
		for i = leftPtr; i < right; i++ {
			leftPtr++
			if i < right && a.array[leftPtr] < pivot {
				continue
			} else {
				break
			}
		}

		for i = rightPtr; i > left; i-- {
			rightPtr--
			if i > left && a.array[rightPtr] >= pivot {
				continue
			} else {
				break
			}
		}

		if leftPtr >= rightPtr {
			break
		} else {
			temp := a.array[leftPtr]
			a.array[leftPtr] = a.array[rightPtr]
			a.array[rightPtr] = temp
		}
	}
	return leftPtr
}

func (a *PartArray) Display() {
	for _, v := range a.array {
		fmt.Print(v, "; ")
	}
	fmt.Println()
}
