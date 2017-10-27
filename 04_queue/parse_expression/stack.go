package parse_expression

type operatorsStack struct {
	data    []string
	maxSize int
	top     int
}

func NewOperandsStack(size int) *operatorsStack {
	return &operatorsStack{
		data:    make([]string, size),
		maxSize: size,
		top:     -1,
	}
}

func (s *operatorsStack) push(char string) {
	s.top++
	s.data[s.top] = char
}

func (s *operatorsStack) pop() string {
	char := s.data[s.top]
	s.top--
	return char
}

func (s *operatorsStack) peek() string {
	return s.data[s.top]
}

func (s *operatorsStack) peekN(n int) string {
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
