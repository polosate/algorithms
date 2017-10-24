package parse_expression

import "fmt"

type inToPost struct {
	stack   *operatorsStack
	inBytes []byte
	outStr  []byte
}

func NewInToPost(inStr string) *inToPost {
	inputBytes := []byte(inStr)
	return &inToPost{
		stack:   NewOperandsStack(len(inputBytes)),
		inBytes: inputBytes,
		outStr:  []byte{},
	}
}

// A + B × (C – D)
// 1.

func (this *inToPost) DoTrans() {
	for _, v := range this.inBytes {
		switch v {
		case string("+"), string("-"):
			this.gotOp(v, 1)
		case string("*"), string("/"):
			this.gotOp(v, 2)
		case string("("):
			this.stack.push(v)
		case string(")"):
			fmt.Println()
		default:
			this.outStr = append(this.outStr, v)
		}
	}
}

func (this *inToPost) gotOp(op byte, priority1 int) {
	for !this.stack.isEmpty() {
		stOp := this.stack.pop()
		if stOp == byte("(") {
			this.stack.push(stOp)
			break
		} else {
			var priority2 int
			switch stOp {
			case byte("+"), byte("-"):
				priority2 = 1
			case byte("*"), byte("/"):
				priority2 = 2
			}
			if priority2 < priority1 {
				this.stack.push(stOp)
				break
			} else {
				this.outStr = append(this.outStr, stOp)
			}
		}
	}
	this.stack.push(op)
}
