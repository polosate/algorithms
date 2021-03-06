package selecting

type SArray struct {
	data []int64
	len  int
}

// For benchmark
func NewSArray(array []int64) SArray {
	return SArray{
		data: array,
		len:  len(array),
	}
}

func NewArray(size int) SArray {
	return SArray{make([]int64, size), 0}
}

func (this *SArray) Insert(elem int64) bool {
	if this.len >= cap(this.data) {
		return false
	}
	this.data[this.len] = elem
	this.len++
	return true
}

func (this *SArray) SelectingSort() {
	var temp int64
	var minInd int
	for i := 0; i < this.len-1; i++ {
		minInd = i
		for j := i + 1; j < this.len; j++ {
			if this.data[j] < this.data[minInd] {
				minInd = j
			}
		}
		temp = this.data[i]
		this.data[i] = this.data[minInd]
		this.data[minInd] = temp
	}
}

func (this *SArray) GetData() []int64 {
	return this.data[:this.len]
}

func (this *SArray) GetLen() int {
	return this.len
}
