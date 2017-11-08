package _7_recursion

import "fmt"

type anagram struct {
	chArr []byte
	size  int
}

func NewAnagram(word string) anagram {
	return anagram{
		chArr: []byte(word),
		size:  len([]byte(word)),
	}
}

func (a *anagram) Do() {
	a.doAnagram(a.size)
}

func (a *anagram) doAnagram(newSize int) {
	if newSize == 1 {
		return
	}
	for i := 0; i < newSize; i++ {
		a.doAnagram(newSize - 1)
		if newSize == 2 {
			a.display()
		}
		a.shift(newSize)
	}
}

func (a *anagram) display() {
	for i := 0; i < a.size; i++ {
		fmt.Print(string(a.chArr[i]))
	}
	fmt.Println()
}

func (a *anagram) shift(newSize int) {
	position := a.size - newSize
	temp := a.chArr[position]
	for i := position + 1; i < a.size; i++ {
		a.chArr[i-1] = a.chArr[i]
	}
	a.chArr[a.size-1] = temp
}
