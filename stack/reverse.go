package stack

type charStack struct {
	bArray []byte
	top    int
	size   int
}

func newCharStack(size int) charStack {
	return charStack{
		bArray: make([]byte, size),
		top: -1,
		size: size,
	}
}
