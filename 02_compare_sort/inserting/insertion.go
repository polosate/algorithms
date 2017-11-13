package inserting

type IArray struct {
	data []int64
	len  int
}

// For benchmark
func NewIArray(array []int64) IArray {
	return IArray{
		data: array,
		len:  len(array),
	}
}

func NewArray(size int) IArray {
	return IArray{make([]int64, size), 0}
}

func (this *IArray) Insert(elem int64) bool {
	if this.len >= cap(this.data) {
		return false
	}
	this.data[this.len] = elem
	this.len++
	return true
}

func (this *IArray) Median() float64 {
	this.InsertingSort()
	if this.len%2 == 0 {
		return float64(this.data[this.len/2]+this.data[this.len/2-1]) / 2
	} else {
		return float64(this.data[this.len/2])
	}
}

func (this *IArray) NoDups() {
	this.InsertingSort()
	insertInd := 0
	for i := 0; i < this.len-1; i++ {
		if this.data[i] != this.data[i+1] {
			insertInd++
			this.data[insertInd] = this.data[i+1]
		}
	}
	this.len = insertInd + 1
}

func (this *IArray) InsertingSortWithBinarySearch() {
	var temp int64

	for i := 1; i < this.len; i++ {
		temp = this.data[i]
		insertIn := binarySearch(this.data[0:i], temp)
		for j := insertIn; j < i; j++ {
			this.data[j+1] = this.data[j]
		}
		this.data[insertIn] = temp
	}
}

func (this *IArray) InsertingSort() {
	var tempIn int
	var temp int64
	for i := 1; i < this.len; i++ {
		tempIn = i
		temp = this.data[tempIn]
		for j := i - 1; j >= 0; j-- {
			if this.data[j] < temp {
				break
			}
			this.data[j+1] = this.data[j]
			tempIn = j
		}
		this.data[tempIn] = temp
	}
}

func (this *IArray) GetData() []int64 {
	return this.data[:this.len]
}

func (this *IArray) GetLen() int {
	return this.len
}

func binarySearch(array []int64, el int64) int {
	lowerBound := 0
	upperBound := len(array) - 1
	for {
		curIn := (lowerBound + upperBound) / 2
		if el > array[curIn] && (len(array) == 1 || el < array[curIn+1]) {
			curIn++
			return curIn
		} else if el < array[curIn] && (len(array) == 1 || el > array[curIn-1]) {
			curIn--
			return curIn
		} else if el > array[curIn] {
			lowerBound = curIn + 1
		} else if el < array[curIn] {
			upperBound = curIn - 1
		}
	}
}
