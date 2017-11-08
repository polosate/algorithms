package matrix

import (
	"errors"
	"fmt"
	"log"
)

type matrix struct {
	firstList *list
	x         int
	y         int
}

func NewMatrix(x, y int) *matrix {
	newMatrix := matrix{
		x: x,
		y: y,
	}
	// create first list
	firstList := newMatrix.emptyList()
	newMatrix.firstList = firstList

	// create other lists
	currentList := newMatrix.firstList
	for i := 0; i < x-1; i++ {
		list := newMatrix.emptyList()
		currentList.nextList = list
		currentList = list
	}
	return &newMatrix
}

func (m *matrix) SetValue(x, y int, value float32) {
	if x > m.x-1 || y-1 > m.y-1 {
		log.Println("incorrect coordinates")
		return
	}
	currentList := m.firstList
	for i := 0; i < x; i++ {
		currentList = currentList.nextList
	}
	currentItem := currentList.first
	for i := 0; i < y; i++ {
		currentItem = currentItem.next
	}
	currentItem.SetValue(value)
}

func (m *matrix) GetValue(x, y int) (float32, error) {
	if x > m.x-1 || y-1 > m.y-1 {
		return -1, errors.New("incorrect coordinates")
	}
	currentList := m.firstList
	for i := 0; i < x; i++ {
		currentList = currentList.nextList
	}
	currentItem := currentList.first
	for i := 0; i < y; i++ {
		currentItem = currentItem.next
	}
	return currentItem.GetValue(), nil
}

func (m *matrix) DisplayMatrix() {
	currentList := m.firstList
	for currentList != nil {
		currentList.DisplayList()
		currentList = currentList.nextList
	}
	fmt.Println()
}

func (m *matrix) emptyList() *list {
	firstList := NewList()
	iter := firstList.GetIterator()
	for j := 0; j < m.y; j++ {
		iter.InsertAfter(0)
	}
	return firstList
}
