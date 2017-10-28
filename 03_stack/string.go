package stack

type StrStack struct {
	abstractStack
}

func NewStrStack(size int) StrStack {
	return StrStack{
		newAbstractStack(size),
	}
}

func (s *StrStack) Push(char string) {
	s.abstractStack.push(char)
}

func (s *StrStack) Pop() string {
	res, err := s.abstractStack.pop()
	if err != nil {
		return ""
	}
	char, _ := res.(string)

	return char
}

func (s *StrStack) Peek() string {
	res, err := s.abstractStack.peek()
	if err != nil {
		return ""
	}
	char, _ := res.(string)

	return char
}

func (s *StrStack) IsEmpty() bool {
	return s.abstractStack.isEmpty()
}

func (s *StrStack) isFull() bool {
	return s.abstractStack.isFull()
}

func (s *StrStack) size() int {
	return s.abstractStack.top + 1
}
