package _7_recursion

import "fmt"

type MArray struct {
	array []int64
	size  int
}

func NewMArray(array []int64) MArray {
	return MArray{
		array: array,
		size:  len(array),
	}
}

func (ma *MArray) MergeSort() {
	workArea := make([]int64, ma.size)
	ma.recMergeSort(workArea, 0, ma.size-1)
}

func (ma *MArray) recMergeSort(workArea []int64, lowerBound, upperBound int) {
	if lowerBound == upperBound {
		return
	} else {
		mid := (lowerBound + upperBound) / 2
		ma.recMergeSort(workArea, lowerBound, mid)
		ma.recMergeSort(workArea, mid+1, upperBound)
		ma.merge(workArea, lowerBound, mid, upperBound)
	}
}

func (ma *MArray) merge(workArea []int64, lowerBound, mid, upperBound int) {
	j := 0
	aInd := lowerBound
	bInd := mid + 1
	for aInd <= mid && bInd <= upperBound {
		if ma.array[aInd] < ma.array[bInd] {
			workArea[j] = ma.array[aInd]
			aInd++
		} else {
			workArea[j] = ma.array[bInd]
			bInd++
		}
		j++
	}
	for aInd <= mid {
		workArea[j] = ma.array[aInd]
		aInd++
		j++
	}
	for bInd <= upperBound {
		workArea[j] = ma.array[bInd]
		bInd++
		j++
	}
	for j = 0; j <= upperBound-lowerBound; j++ {
		ma.array[lowerBound+j] = workArea[j]
	}
}

func (ma *MArray) Display() {
	fmt.Println(ma.array)
}

func mergeSort(arrayA, arrayB []int) []int {
	sizeA := len(arrayA)
	sizeB := len(arrayB)
	arrayC := make([]int, 0, sizeA+sizeB)

	var a, b int
	a = 0
	b = 0

	for a < sizeA && b < sizeB {
		if arrayA[a] < arrayB[b] {
			arrayC = append(arrayC, arrayA[a])
			a++
		} else {
			arrayC = append(arrayC, arrayB[b])
			b++
		}
	}
	for a < sizeA {
		arrayC = append(arrayC, arrayA[a])
		a++
	}
	for b < sizeB {
		arrayC = append(arrayC, arrayB[b])
		b++
	}
	return arrayC
}
