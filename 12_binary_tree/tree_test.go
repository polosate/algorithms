package _2_binary_tree

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := NewTree()

	tree.Insert(55)
	current, parent := tree.find(55)

	//for i := 0; i < 10; i++ {
	//	tree.Insert(int64(rand.Intn(50)))
	//}
	//tree.TraversInOrder()
}
