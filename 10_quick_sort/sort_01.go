package _0_quick_sort

import "fmt"

type QuickSortArray struct {
	array []int64
}

func NewQuickSortArray(array []int64) QuickSortArray {
	return QuickSortArray{array: array}
}

func (a *QuickSortArray) QuickSort() {
	a.reqQuickSort(0, len(a.array)-1)
}

func (a *QuickSortArray) reqQuickSort(left, right int) {
	if right-left <= 0 {
		return
	} else {
		pivot := a.array[right]
		partition := a.partition(left, right, pivot)
		a.reqQuickSort(left, partition-1)
		a.reqQuickSort(partition+1, right)
	}
}

func (a *QuickSortArray) partition(left, right int, pivot int64) int {
	leftPtr := left - 1
	rightPtr := right
	i := 0
	for {
		for i = leftPtr; i < right; i++ {
			leftPtr++
			if a.array[leftPtr] < pivot {
				continue
			} else {
				break
			}
		}

		for i = rightPtr; i > left; i++ {
			rightPtr--
			if a.array[rightPtr] > pivot {
				continue
			} else {
				break
			}
		}

		if leftPtr >= rightPtr {
			break
		} else {
			a.swap(leftPtr, rightPtr)
		}
	}
	a.swap(leftPtr, right)
	return leftPtr
}

func (a *QuickSortArray) swap(ind1, ind2 int) {
	temp := a.array[ind1]
	a.array[ind1] = a.array[ind2]
	a.array[ind2] = temp
}

func (a *QuickSortArray) Display() {
	for i := range a.array {
		fmt.Print(a.array[i], "; ")
	}
	fmt.Println()
}
