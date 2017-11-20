package _9_partition

func Partition(array []int64) {
	recPartition(array, 0, len(array)-1, len(array)-1)
}

func recPartition(array []int64, left, right, pivotInd int) {
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
}

func swap(array []int64, ind1, ind2 int) {
	temp := array[ind1]
	array[ind1] = array[ind2]
	array[ind2] = temp
}
