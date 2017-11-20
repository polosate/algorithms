package _9_partition

func Partition(array []int64) {
	partition(array, 0, len(array)-1, len(array)-1)
}

func Median(array []int64) int64 {
	mid := len(array) / 2
	medInd := median(array, mid, 0, len(array)-1)
	return array[medInd]
}

func median(array []int64, mid int, left, right int) int {
	ind := partition(array, left, right, right)
	if ind == mid {
		return ind
	}
	if ind > mid {
		return median(array, mid, left, ind-1)
	} else {
		return median(array, mid, ind, right)
	}
}

func partition(array []int64, left, right, pivotInd int) int {
	leftPtr := left - 1
	rightPtr := right

	for {
		for leftPtr < right {
			leftPtr++
			if array[leftPtr] < array[pivotInd] {
				continue
			} else {
				break
			}
		}

		for rightPtr > left {
			rightPtr--
			if array[rightPtr] > array[pivotInd] {
				continue
			} else {
				break
			}
		}
		if leftPtr >= rightPtr {
			break
		} else {
			swap(array, leftPtr, rightPtr)
		}
	}
	swap(array, leftPtr, right)
	return leftPtr
}

func swap(array []int64, ind1, ind2 int) {
	temp := array[ind1]
	array[ind1] = array[ind2]
	array[ind2] = temp
}
