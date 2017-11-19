package _7_recursion

import (
	"reflect"
	"testing"
)

func TestKnuthSeq(t *testing.T) {
	res := knuthSeq(100)
	expected := []int{1, 4, 13, 40}
	if !reflect.DeepEqual(res, expected) {
		t.Error("want sequence", expected, "have", res)
	}
}
