package parse_expression

import (
	s "data_structures_and_algorithms/03_stack"
)

type inToPost struct {
	stack  s.StrStack
	inStr  string
	outStr string
}

func NewInToPost(inStr string) *inToPost {
	return &inToPost{
		stack:  s.NewStrStack(len(inStr)),
		inStr:  inStr,
		outStr: "",
	}
}

// A+B×(C–D)
// ABCD-*+
func (this *inToPost) DoTrans() {
	for _, char := range this.inStr {
		switch string(char) {
		case string('+'), string('-'):
			this.gotOp(string(char), 1)
		case string('*'), string('/'):
			this.gotOp(string(char), 2)
		case string('('):
			this.stack.Push(string(char))
		case string(')'):
			this.gotParen()
		default:
			this.outStr += string(char)
		}
	}
	for !this.stack.IsEmpty() {
		this.outStr += this.stack.Pop()
	}
}

func (this *inToPost) gotOp(op string, priority1 int) {
	for !this.stack.IsEmpty() {
		stOp := this.stack.Pop()
		if stOp == string('(') {
			this.stack.Push(stOp)
			break
		} else {
			var priority2 int
			switch stOp {
			case string('+'), string('-'):
				priority2 = 1
			case string('*'), string('/'):
				priority2 = 2
			}
			if priority2 < priority1 {
				this.stack.Push(stOp)
				break
			} else {
				this.outStr += stOp
			}
		}
	}
	this.stack.Push(op)
}

func (this *inToPost) gotParen() {
	for !this.stack.IsEmpty() {
		el := this.stack.Pop()
		if el == string('(') {
			break
		} else {
			this.outStr += el
		}
	}
}
