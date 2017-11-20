package _9_partition

import (
	"testing"
)

func TestMedian(t *testing.T) {
	array := []int64{0, 3, 1, 6, 4, 5, 2, 7}
	med := Median(array)
	if med != 4 {
		t.Error("Want median", 4, "have", med)
	}

	array = []int64{0, 3, 1, 6, 4, 5, 2}
	med = Median(array)
	if med != 3 {
		t.Error("Want median", 3, "have", med)
	}

	array = []int64{0, 3, 1, 6, 4, 5, 7}
	med = Median(array)
	if med != 4 {
		t.Error("Want median", 4, "have", med)
	}

	array = []int64{0, 3, 1, 6, 4, 5, 7, 2}
	med = Median(array)
	if med != 4 {
		t.Error("Want median", 4, "have", med)
	}
}

//
//func TestSmallPartition(t *testing.T) {
//
//	seed := []int64{1, 5, 3}
//	combi := doCombinations(seed, []int64{})
//
//	for _, v := range combi {
//		fmt.Println("Before: ", v)
//		Partition(v)
//		fmt.Println("After: ", v)
//	}
//
//	seed = []int64{1, 2}
//	combi = doCombinations(seed, []int64{})
//
//	for _, v := range combi {
//		fmt.Println("Before: ", v)
//		Partition(v)
//		fmt.Println("After: ", v)
//	}
//
//	v := []int64{5}
//	fmt.Println("Before: ", v)
//	Partition(v)
//	fmt.Println("After: ", v)
//
//	v = []int64{5, 7, 2, 3}
//	fmt.Println("Before: ", v)
//	Partition(v)
//	fmt.Println("After: ", v)
//}
//
//func doCombinations(array []int64, prefix []int64) [][]int64 {
//	if len(array) == 0 {
//		res := [][]int64{}
//		res = append(res, prefix)
//		return res
//	}
//	var res [][]int64
//	for i := 0; i < len(array); i++ {
//		r := doCombinations(array[1:], append(prefix, array[0]))
//		res = append(res, r...)
//		shift(array)
//	}
//	return res
//}
//
//func shift(array []int64) {
//	temp := array[len(array)-1]
//	for i := len(array) - 2; i >= 0; i-- {
//		array[i+1] = array[i]
//	}
//	array[0] = temp
//}
