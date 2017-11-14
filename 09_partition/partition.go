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
	leftPtr := left
	rightPtr := right

	for {
		for leftPtr < right && a.array[leftPtr] < pivot {
			leftPtr++
		}

		for rightPtr > left && a.array[rightPtr] > pivot {
			rightPtr--
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
