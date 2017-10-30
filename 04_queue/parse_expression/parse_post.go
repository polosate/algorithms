package parse_expression

import (
	s "data_structures_and_algorithms/03_stack"
	"fmt"
)

type parsePost struct {
	stack s.DoubleStack
	inStr string
}

func NewParsePost(input string) parsePost {
	return parsePost{
		stack: s.NewDoubleStack(len(input)),
		inStr: input,
	}
}

// 1243-*+
// 1243
//
func (pp *parsePost) DoCalculate() float64 {
	var num1, num2, res float64
	for _, char := range pp.inStr {
		switch string(char) {
		case string("+"), string("-"), string("*"), string("/"):
			num1, _ = pp.stack.Pop()
			num2, _ = pp.stack.Pop()
			switch string(char) {
			case string("+"):
				res = num2 + num1
			case string("-"):
				res = num2 - num1
			case string("*"):
				res = num2 * num1
			case string("/"):
				res = num2 / num1
			}
			pp.stack.Push(res)
		default:
			pp.stack.Push(float64(char))
		}
		fmt.Println(string(char))
		fmt.Println(pp.stack)
		fmt.Println("---------")
	}
	res, _ = pp.stack.Pop()
	return res
}
