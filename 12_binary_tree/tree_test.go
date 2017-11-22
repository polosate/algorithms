package _2_binary_tree

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
