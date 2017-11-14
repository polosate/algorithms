package shell_sort

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	array := []int64{5, 2, 7, 1, 9, 8, 3, 6, 4, 0}
	shell := NewShellArray(array)
	shell.ShellSort()
	expected := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(shell.array, expected) {
		t.Error("Want", expected, "have", shell.array)
	}
}
