package exercises

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tr := NewTree([]string{"A", "B", "C"})
	fmt.Println(tr.root.value)
	fmt.Println(tr.root.leftChild.value)
	fmt.Println(tr.root.rightChild.value)
	fmt.Println(tr.root.leftChild.leftChild.value)
	fmt.Println(tr.root.leftChild.rightChild.value)
}

//          +
//    +           +
// +     +     +    g
//a b   c d   e f
//
//
//+
//++
//+++a
//gfedcb
func TestBalancedTree(t *testing.T) {
	NewBalancedTree([]string{"A", "B", "C", "D", "E", "F", "G"})
	//fmt.Println(tr.root.value)
	//fmt.Print(tr.root.leftChild.value)
	//fmt.Println(tr.root.rightChild.value)
	//fmt.Print(tr.root.leftChild.leftChild.value)
	//fmt.Print(tr.root.leftChild.rightChild.value)
	//fmt.Print(tr.root.rightChild.leftChild.value)
	//fmt.Println(tr.root.rightChild.rightChild.value)
}
