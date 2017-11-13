package bubble

type BArray struct {
	data []int64
	len  int
}

// For benchmark
func NewBArray(array []int64) BArray {
	return BArray{
		data: array,
		len:  len(array),
	}
}

func NewArray(size int) BArray {
	return BArray{make([]int64, size), 0}
}

func (this *BArray) Insert(elem int64) bool {
	if this.len >= cap(this.data) {
		return false
	}
	this.data[this.len] = elem
	this.len++
	return true
}

func (this *BArray) BubbleSort() {
	var temp int64
	for i := 0; i < this.len-1; i++ {
		for j := 0; j < this.len-i-1; j++ {
			if this.data[j] > this.data[j+1] {
				temp = this.data[j]
				this.data[j] = this.data[j+1]
				this.data[j+1] = temp
			}
		}
	}
}

func (this *BArray) GetData() []int64 {
	return this.data[:this.len]
}

func (this *BArray) GetLen() int {
	return this.len
}
