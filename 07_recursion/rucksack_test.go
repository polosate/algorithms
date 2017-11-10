package _7_recursion

import (
	"reflect"
	"testing"
)

func TestRucksackPass(t *testing.T) {
	rucksack := NewRucksack(20, []int{11, 8, 7, 6, 5})
	expected := []int{8, 7, 5}
	result := rucksack.Do()
	if !reflect.DeepEqual(result, expected) {
		t.Error("Want", expected, "have", result)
	}
}

func TestRucksackFail(t *testing.T) {
	rucksack := NewRucksack(22, []int{11, 8, 7, 6, 5})
	expected := []int{0}
	result := rucksack.Do()
	if !reflect.DeepEqual(result, expected) {
		t.Error("Want", expected, "have", result)
	}
}
