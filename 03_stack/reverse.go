package stack

// Reverse returns reverted string
func Reverse(s string) (string, error) {
	byteWord := []byte(s)
	stack := NewByteStack(len(byteWord))

	for _, v := range byteWord {
		stack.Push(v)
	}

	i := 0
	for !stack.isEmpty() {
		if b, err := stack.Pop(); err != nil {
			return "", err
		} else {
			byteWord[i] = b
			i++
		}
	}
	return string(byteWord), nil
}
