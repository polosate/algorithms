package stack

import (
	"errors"
	"fmt"
)

// BracketsChecker checks that brackets in input string has right order
func BracketsChecker(s string) error {
	byteS := []byte(s)

	stack := NewByteStack(len(byteS))

	for i := 0; i < len(byteS); i++ {
		inputChar := byteS[i]
		switch string(inputChar) {
		case "{", "(", "[":
			stack.Push(inputChar)
		case "}", ")", "]":
			stackChar, err := stack.Pop()
			if err != nil {
				return errors.New(fmt.Sprintf("pop error: %s", err.Error()))
			}
			if (string(inputChar) == "}" && string(stackChar) != "{") ||
				(string(inputChar) == ")" && string(stackChar) != "(") ||
				(string(inputChar) == "]" && string(stackChar) != "[") {
				return errors.New(fmt.Sprintf("expected %s, but received %s in %d",
					reverseBracket(stackChar), string(inputChar), i))
			}
		default:
			break
		}
	}
	if !stack.IsEmpty() {
		return errors.New("not enough closing brackets")
	}
	return nil
}

func reverseBracket(char byte) string {
	switch string(char) {
	case "{":
		return "}"
	case "(":
		return ")"
	case "[":
		return "]"
	default:
		return "xxx"
	}
}
