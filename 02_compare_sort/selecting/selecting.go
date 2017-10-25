package selecting

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

func (this *array) SelectingSort() {
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

func (this *array) GetData() []int64 {
	return this.data[:this.len]
}

func (this *array) GetLen() int {
	return this.len
}
