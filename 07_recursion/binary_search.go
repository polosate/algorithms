package _7_recursion

type array struct {
	intArr []int
	size   int
}

func NewArray(arr []int) array {
	return array{
		intArr: arr,
		size:   len(arr),
	}
}

func (a *array) Find(key int) int {
	return a.recSearch(0, a.size-1, key)
}

func (a *array) recSearch(lowerBound, upperBound int, searchKey int) int {
	idx := (lowerBound + upperBound) / 2
	if lowerBound > upperBound {
		return a.size
	}
	if a.intArr[idx] == searchKey {
		return idx
	}

	if a.intArr[idx] > searchKey {
		return a.recSearch(lowerBound, idx-1, searchKey)
	} else {
		return a.recSearch(idx+1, upperBound, searchKey)
	}
}
