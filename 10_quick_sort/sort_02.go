package _0_quick_sort

import "fmt"

type QuickSortArray2 struct {
	array []int64
	level int
}

func NewQuickSortArray2(array []int64) QuickSortArray2 {
	return QuickSortArray2{array: array, level: 0}
}

func (a *QuickSortArray2) QuickSort() {
	a.reqQuickSort(0, len(a.array)-1)
}

func (a *QuickSortArray2) reqQuickSort(left, right int) {
	if right-left+1 <= 3 {
		a.manualSort(left, right)
	} else {
		median := a.medianOf3(left, right)
		partition := a.partition(left, right, median)
		a.reqQuickSort(left, partition-1)
		a.reqQuickSort(partition+1, right)
	}
}

func (a *QuickSortArray2) medianOf3(left, right int) int64 {
	center := (left + right) / 2
	if a.array[left] > a.array[center] {
		a.swap(left, center)
	}
	if a.array[left] > a.array[right] {
		a.swap(left, right)
	}
	if a.array[center] > a.array[right] {
		a.swap(center, right)
	}
	a.swap(center, right-1)
	return a.array[right-1]
}

func (a *QuickSortArray2) partition(left, right int, pivot int64) int {
	leftPtr := left
	rightPtr := right - 1
	for {
		for {
			leftPtr++
			if a.array[leftPtr] < pivot {
				continue
			} else {
				break
			}
		}

		for {
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
	a.swap(leftPtr, right-1)
	return leftPtr
}

func (a *QuickSortArray2) manualSort(left, right int) {
	size := right - left + 1
	if size <= 1 {

	} else if size == 2 {
		if a.array[left] > a.array[right] {
			a.swap(left, right)
		}

	} else if size == 3 {
		if a.array[left] > a.array[right-1] {
			a.swap(left, right-1)
		}
		if a.array[left] > a.array[right] {
			a.swap(left, right)
		}
		if a.array[right-1] > a.array[right] {
			a.swap(right-1, right)
		}
	}
}

func (a *QuickSortArray2) swap(ind1, ind2 int) {
	temp := a.array[ind1]
	a.array[ind1] = a.array[ind2]
	a.array[ind2] = temp
}

func (a *QuickSortArray2) Display() {
	for i := range a.array {
		fmt.Print(a.array[i], "; ")
	}
	fmt.Println()
}
