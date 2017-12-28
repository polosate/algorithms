package exercises

import "github.com/pkg/errors"

type stacksSet struct {
	size   int
	stacks []iStack
}

func newStacksSet(size int) stacksSet {
	return stacksSet{}
}

func (s *stacksSet) push(data int64) {
	curStack := s.getLastStack()
	if curStack == nil || curStack.isFull() {
		newStack := newStack(s.size)
		newStack.push(data)
		s.stacks = append(s.stacks, newStack)
	} else {
		curStack.push(data)
	}
}

func (s *stacksSet) pop() (int64, error) {
	curStack := s.getLastStack()
	if curStack == nil {
		return 0, errors.New("stack is empty")
	}
	data, err := curStack.pop()
	if err != nil {
		return 0, err
	}
	if curStack.isEmpty() {
		s.stacks[len(s.stacks)-1] = nil
	}
	return data, nil
}

func (s *stacksSet) getLastStack() iStack {
	if len(s.stacks) == 0 {
		return nil
	}
	return s.stacks[len(s.stacks)-1]
}
