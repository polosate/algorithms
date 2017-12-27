package exercises

import "testing"

func TestIsPalindromeOdd_True(t *testing.T) {
	list := NewSingleList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(2)
	list.InsertFirst(1)

	if !isPalindrome(list) {
		t.Error()
	}
}

func TestIsPalindromeEven_True(t *testing.T) {
	list := NewSingleList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(3)
	list.InsertFirst(2)
	list.InsertFirst(1)

	if !isPalindrome(list) {
		t.Error()
	}
}

func TestIsPalindromeOdd_False(t *testing.T) {
	list := NewSingleList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(2)
	list.InsertFirst(6)

	if isPalindrome(list) {
		t.Error()
	}
}

func TestIsPalindromeEven_False(t *testing.T) {
	list := NewSingleList()
	list.InsertFirst(1)
	list.InsertFirst(2)
	list.InsertFirst(3)
	list.InsertFirst(3)
	list.InsertFirst(4)
	list.InsertFirst(1)

	if isPalindrome(list) {
		t.Error()
	}
}
