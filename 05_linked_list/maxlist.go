package main

import (
	"data_structures_and_algorithms/05_linked_list/single"
	"fmt"
)

// Lenovo ThinkPad X1 Carbon Gen 1 8GB
// 236.087.000 (~4GB)

// Lenovo ThinkPad T450 16GB
// 621.780.000 (~12GB)
func main() {
	list := single.NewLinkList()
	i := 0
	for {
		if i%1000 == 0 {
			fmt.Println(i)
		}
		list.InsertFirst(1, 1)
		i++
	}
}
