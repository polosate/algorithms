package _2_binary_tree

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := NewTree()

	tree.Insert(55)
	current, parent := tree.find(55)

	if current.GetValue() != 55 {
		t.Error("Want", 55, "have", current.GetValue())
	}
	if parent != nil {
		t.Error("Want", nil, "have", parent)
	}

	tree.Insert(22)
	current, parent = tree.find(22)
	if current.GetValue() != 22 {
		t.Error("Want", 22, "have", current.GetValue())
	}
	if parent.GetValue() != 55 {
		t.Error("Want", 55, "have", parent.GetValue())
	}
	if parent.leftChild != current {
		t.Error("Warning!")
	}

	//for i := 0; i < 10; i++ {
	//	tree.Insert(int64(rand.Intn(50)))
	//}
	//tree.TraversInOrder()
}
