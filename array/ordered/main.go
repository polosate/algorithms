package ordered

type array struct {
	data []int64
	len  int
}

func newArray(size int) array {
	return array{make([]int64, size), 0}
}

func (this *array) Insert(elem int64) bool {
	if this.len >= cap(this.data) {
		return false
	}

	var i int
	for i = 0; i < this.len; i++ {
		if this.data[i] > elem {
			break
		}
	}
	for j := this.len - 1; j >= i; j-- {
		this.data[j+1] = this.data[j]
	}

	this.data[i] = elem
	this.len++
	return true
}

func (this *array) LinearSearch(elem int64) int {
	for i := 0; i < this.len; i++ {
		if this.data[i] == elem {
			return i
		} else if this.data[i] > elem {
			break
		}
	}
	return -1
}

func (this *array) BinarySearch(elem int64) int {
	var currentIn int
	lowerBound := 0
	upperBound := this.len - 1

	for {
		currentIn = (lowerBound + upperBound) / 2
		if this.data[currentIn] == elem {
			return currentIn
		} else if lowerBound > upperBound {
			return -1
		} else if this.data[currentIn] < elem {
			lowerBound = currentIn + 1
		} else if this.data[currentIn] > elem {
			upperBound = currentIn - 1
		}
	}
}

func (this *array) Delete(elem int64) bool {
	i := this.BinarySearch(elem)
	if i == -1 {
		return false
	}
	this.len--
	for j := i; j < this.len; j++ {
		this.data[j] = this.data[j+1]
	}
	return true
}

func (this *array) Reverse() {
	var temp int64
	for i := 0; i < this.len/2; i++ {
		temp = this.data[i]
		this.data[i] = this.data[this.len-i-1]
		this.data[this.len-i-1] = temp
	}
}

func (this *array) GetData() []int64 {
	return this.data[:this.len]
}

func (this *array) GetLen() int {
	return this.len
}
