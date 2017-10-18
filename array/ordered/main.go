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

	

	for j := this.len; j >= i; j-- {
		this.data[j+1] = this.data[j]
	}

	this.data[i] = elem
	this.len++
	return true
}

func (this *array) Find(elem int64) int {
	for i := 0; i < this.len; i++ {
		if this.data[i] == elem {
			return i
		} else if this.data[i] > elem {
			break
		}
	}
	return -1
}

func (this *array) Delete(elem int64) bool {
	i := this.Find(elem)
	if i == -1 {
		return false
	}
	this.len--
	for j := i; j < this.len; j++ {
		this.data[j] = this.data[j+1]
	}
	return true
}

func (this *array) GetData() []int64 {
	return this.data[:this.len]
}

func (this *array) GetLen() int {
	return this.len
}
