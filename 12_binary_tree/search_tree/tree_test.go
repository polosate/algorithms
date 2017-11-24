package search_tree

import (
	"testing"
)

func TestTreeInsert(t *testing.T) {
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
}

func TestTreeRemoveRootWithoutChildren(t *testing.T) {
	tree := NewTree()

	tree.Insert(55)
	el := tree.Remove(55)
	if el.GetValue() != 55 {
		t.Error()
	}
	if tree.root != nil {
		t.Error()
	}
}

func TestTreeRemoveRootWithLeftChild(t *testing.T) {
	tree := NewTree()

	tree.Insert(60)
	tree.Insert(50)
	el := tree.Remove(60)
	if el.GetValue() != 60 {
		t.Error()
	}
	if tree.root.GetValue() != 50 {
		t.Error()
	}
}

func TestTreeRemoveRootWithRightChild(t *testing.T) {
	tree := NewTree()

	tree.Insert(60)
	tree.Insert(70)
	el := tree.Remove(60)
	if el.GetValue() != 60 {
		t.Error()
	}
	if tree.root.GetValue() != 70 {
		t.Error()
	}
}

func TestNewTreeRemoveRootWithBothChildrenSuccessorIsRightChild(t *testing.T) {
	tree := NewTree()

	tree.Insert(60)
	tree.Insert(70)
	tree.Insert(50)
	tree.Insert(80)
	el := tree.Remove(60)
	if el.GetValue() != 60 {
		t.Error()
	}
	if tree.root.GetValue() != 70 {
		t.Error()
	}
	if tree.root.leftChild.GetValue() != 50 {
		t.Error()
	}
	if tree.root.rightChild.GetValue() != 80 {
		t.Error()
	}
}

func TestNewTreeRemoveRootWithBothChildrenSuccessorIsLeftChild(t *testing.T) {
	tree := NewTree()
	tree.Insert(60)
	tree.Insert(70)
	tree.Insert(50)
	tree.Insert(80)
	tree.Insert(65)
	tree.Insert(67)
	el := tree.Remove(60)
	if el.GetValue() != 60 {
		t.Error()
	}
	if tree.root.GetValue() != 65 {
		t.Error()
	}
	if tree.root.leftChild.GetValue() != 50 {
		t.Error()
	}
	if tree.root.rightChild.GetValue() != 70 {
		t.Error()
	}
	if tree.root.rightChild.leftChild.GetValue() != 67 {
		t.Error()
	}
}

func TestTreeRemoveLeaf(t *testing.T) {
	tree := NewTree()

	tree.Insert(61)
	tree.Insert(53)
	el := tree.Remove(53)
	if el.GetValue() != 53 {
		t.Error()
	}
	if tree.root.rightChild != nil || tree.root.leftChild != nil {
		t.Error()
	}
}

func TestTreeRemoveNodeWithLeftChild(t *testing.T) {
	tree := NewTree()
	tree.Insert(60)
	tree.Insert(70)
	tree.Insert(50)
	tree.Insert(80)
	tree.Insert(65)
	tree.Insert(63)
	el := tree.Remove(65)
	if el.GetValue() != 65 {
		t.Error()
	}
	if tree.root.rightChild.leftChild.GetValue() != 63 {
		t.Error()
	}
}

func TestTreeRemoveNodeWithRightChild(t *testing.T) {
	tree := NewTree()
	tree.Insert(60)
	tree.Insert(70)
	tree.Insert(50)
	tree.Insert(80)
	tree.Insert(65)
	tree.Insert(67)
	el := tree.Remove(65)
	if el.GetValue() != 65 {
		t.Error()
	}
	if tree.root.rightChild.leftChild.GetValue() != 67 {
		t.Error()
	}
}

func TestTreeRemoveNodeWithBothChildren(t *testing.T) {
	tree := NewTree()
	tree.Insert(60)
	tree.Insert(70)
	tree.Insert(50)
	tree.Insert(80)
	tree.Insert(65)
	tree.Insert(63)
	tree.Insert(67)
	tree.Insert(64)
	tree.Insert(66)
	tree.Insert(68)
	el := tree.Remove(65)
	if el.GetValue() != 65 {
		t.Error()
	}
	if tree.root.rightChild.leftChild.GetValue() != 66 {
		t.Error()
	}
}
