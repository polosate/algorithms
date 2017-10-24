package parse_expression

type operatorsStack struct {
	data    []byte
	maxSize int
	top     int
}

func NewOperandsStack(size int) *operatorsStack {
	return &operatorsStack{
		data:    make([]byte, size),
		maxSize: size,
		top:     -1,
	}
}

func (s *operatorsStack) push(char byte) {
	s.top++
	s.data[s.top] = char
}

func (s *operatorsStack) pop() byte {
	char := s.data[s.top]
	s.top--
	return char
}

func (s *operatorsStack) peek() byte {
	return s.data[s.top]
}

func (s *operatorsStack) peekN(n int) byte {
	return s.data[n]
}

func (s *operatorsStack) isEmpty() bool {
	return s.top == -1
}

func (s *operatorsStack) isFull() bool {
	return s.top == s.maxSize-1
}

func (s *operatorsStack) size() int {
	return s.top + 1
}
