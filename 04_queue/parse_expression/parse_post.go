package parse_expression

import (
	s "data_structures_and_algorithms/03_stack"
)

type parsePost struct {
	stack s.DoubleStack
	input string
}

func NewParsePost(input string) parsePost {
	return parsePost{
		stack: s.NewDoubleStack(len(input)),
		input: input,
	}
}

func (pp *parsePost) DoCalculate() float64 {
	return 0
}
