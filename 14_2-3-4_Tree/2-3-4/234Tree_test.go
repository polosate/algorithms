package __3_4

import (
	"testing"
)

func Test234Tree(t *testing.T) {
	tt := newTree()
	tt.insert(50)
	tt.insert(40)
	tt.insert(60)
	tt.insert(1)
	tt.insert(30)
	tt.insert(70)
	tt.insert(10)
	tt.insert(15)
	tt.insert(25)
	tt.insert(64)
	tt.insert(78)
	tt.insert(16)

	//tt.displayTree()

	min := tt.min()
	if min != 1 {
		t.Error("Want min", 1, "have", min)
	}

	tt.walk()
}
