package bubble

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
	this.data[this.len] = elem
	this.len++
	return true
}

func (this *array) BubbleSort() {
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

func (this *array) GetData() []int64 {
	return this.data[:this.len]
}

func (this *array) GetLen() int {
	return this.len
}
