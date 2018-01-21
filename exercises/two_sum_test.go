package exercises

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	expected := []int{1, 0}
	res := twoSum(nums, target)
	if !reflect.DeepEqual(res, expected) {
		t.Error("want", expected, "have", res)
	}
}
