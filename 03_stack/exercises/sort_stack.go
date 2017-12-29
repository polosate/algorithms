package exercises

type stack1 struct {
	array []int64
	top   int
	size  int
}

func newStack1(size int) stack1 {
	return stack1{
		array: make([]int64, size),
		top:   -1,
		size:  size,
	}
}

func (s *stack1) push(data int64) {
	if s.top == s.size-1 {
		return
	}
	s.top++
	s.array[s.top] = data

}

func (s *stack1) pop() int64 {
	if s.top == -1 {
		return -1
	}
	data := s.array[s.top]
	s.array[s.top] = 0
	s.top--
	return data
}

func (s *stack1) peek() int64 {
	if s.top == -1 {
		return -1
	}
	return s.array[s.top]
}
func (s *stack1) isEmpty() bool {
	return s.top == -1
}

func sortStack(input stack1) stack1 {
	var temp int64
	tempStack := newStack1(input.size)
	for !input.isEmpty() {
		temp = input.pop()

		for !tempStack.isEmpty() && tempStack.peek() > temp {
			input.push(tempStack.pop())

		}
		tempStack.push(temp)
	}
	return tempStack
}
