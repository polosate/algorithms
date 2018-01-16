package exercises

import (
	"errors"
)

// 4; 8; 5; 3; 2; 1;
func FindFromTail(list *singleList, k int) (int64, error) {
	if list.IsEmpty() {
		return -1, errors.New("list is empty")
	}
	iter1 := list.GetIterator()

	i := 0
	for i < k {
		i++
		iter1.NextLink()
		if iter1.GetCurrent() == nil {
			return -1, errors.New("list has length less than k")
		}
	}

	iter2 := list.GetIterator()
	for iter1.GetCurrent().next != nil {
		iter1.NextLink()
		iter2.NextLink()
	}
	return iter2.GetCurrent().value, nil
}
