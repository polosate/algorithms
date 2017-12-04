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

func TestBalancedTree(t *testing.T) {
	tr := NewBalancedTree([]string{"A", "B", "C", "D", "E", "F", "G"})
	r1 := tr.root
	fmt.Println(r1.value)
	fmt.Print(r1.leftChild.value)
	fmt.Println(r1.rightChild.value)
	r2 := r1.leftChild
	r3 := r1.rightChild
	fmt.Print(r2.leftChild.value)
	fmt.Print(r2.rightChild.value)
	fmt.Print(r3.leftChild.value)
	fmt.Println(r3.rightChild.value)
	r5 := r2.rightChild
	r6 := r3.leftChild
	r7 := r3.rightChild
	fmt.Print("-")
	fmt.Print(r5.leftChild.value)
	fmt.Print(r5.rightChild.value)
	fmt.Print(r6.leftChild.value)
	fmt.Print(r6.rightChild.value)
	fmt.Print(r7.leftChild.value)
	fmt.Println(r7.rightChild.value)
}

func TestFullTree(t *testing.T) {
	tr := NewFullTree([]string{"A", "B", "C", "D", "E", "F"})
	r1 := tr.root
	fmt.Println(r1.value)
	fmt.Print(r1.leftChild.value)
	fmt.Println(r1.rightChild.value)
	r2 := r1.leftChild
	r3 := r1.rightChild
	fmt.Print(r2.leftChild.value)
	fmt.Print(r2.rightChild.value)
	fmt.Print(r3.leftChild.value)
	fmt.Println("-")
}
