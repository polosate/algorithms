package _7_recursion

func recMergeSort(array []int) {
	if len(array) == 1 {

	} else {
		rec
	}
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
