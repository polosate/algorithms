package single

import (
	"fmt"
	"math/rand"
)

func SortArray() {
	array := []*Link{}

	for i := 0; i < 10; i++ {
		d := rand.Intn(10)
		array = append(array, NewLink(d, 0))
	}
	fmt.Println("Unsorted array")
	displayArray(array)

	sortedList := NewSortedList()
	for i := range array {
		sortedList.Insert(array[i].iData, array[i].dData)
	}

	for i := 0; i < 10; i++ {
		el, _ := sortedList.Remove()
		array[i] = el
	}

	fmt.Println("Sorted array")
	displayArray(array)
}

func displayArray(array []*Link) {
	for _, v := range array {
		fmt.Print(v.iData, "; ")
	}
	fmt.Println()
}
