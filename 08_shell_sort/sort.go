package shell_sort

import "fmt"

type ShellArray struct {
	array []int64
}

func NewShellArray(array []int64) ShellArray {
	return ShellArray{
		array: array,
	}
}

func (a *ShellArray) ShellSort() {
	h := 1
	for h < len(a.array)/3 {
		h = h*3 + 1
	}

	for h > 0 {
		j := 0
		for i := h; i < len(a.array); i++ {
			current := a.array[i]
			j = i - h
			for j >= h-1 && a.array[j] > current {
				a.array[j+h] = a.array[j]
				j -= h
			}
			a.array[j+h] = current
		}
		h = (h - 1) / 3
	}
}

func (a *ShellArray) Display() {
	for i := range a.array {
		fmt.Print(a.array[i], "; ")
	}
	fmt.Println()
}
