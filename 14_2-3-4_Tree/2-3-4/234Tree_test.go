package _4_2_3_4_Tree

import "testing"

func Test234Tree(t *testing.T) {
	tt := newTree()
	tt.insert(50)
	tt.insert(40)
	tt.insert(60)
	tt.insert(30)
	tt.insert(70)
	tt.insert(10)
	tt.insert(15)
	tt.insert(25)
	tt.insert(64)
	tt.insert(78)
	tt.insert(16)

	tt.displayTree()
}
