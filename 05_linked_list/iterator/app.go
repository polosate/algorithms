package main

import (
	"fmt"

	. "data_structures_and_algorithms/05_linked_list/iterator/iterator"
)

func main() {
	var (
		command string
		value   float32
	)
	list := NewLinkList()
	iter1 := list.GetIterator()

	fmt.Println("Enter first letter of show, reset, ")
	fmt.Println("next, get, before, after, delete: ")
	for {
		fmt.Scanln(&command)
		switch command {
		case "s":
			if list.IsEmpty() {
				fmt.Println("List is empty")
			} else {
				list.DisplayList()
			}
			break
		case "r":
			iter1.Reset()
		case "n":
			if list.IsEmpty() || iter1.AtEnd() {
				fmt.Println("Can't go to next link")
			} else {
				iter1.NextLink()
			}
			break
		case "g":
			if list.IsEmpty() {
				fmt.Println("List is empty")
			} else {
				fmt.Println("Current item:", iter1.GetCurrent().GetData())
			}
			break
		case "b":
			fmt.Print("Enter value: ")
			fmt.Scanln(&value)
			iter1.InsertBefore(value)
			break
		case "a":
			fmt.Print("Enter value: ")
			fmt.Scanln(&value)
			iter1.InsertAfter(value)
			break
		case "d":
			if list.IsEmpty() {
				fmt.Println("List is empty")
			} else {
				fmt.Println("Deleted item:", iter1.DeleteCurrent().GetData())
			}
			break
		default:
			fmt.Printf("Unknown command %q\n", command)
		}
	}
}
