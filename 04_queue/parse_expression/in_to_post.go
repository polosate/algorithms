package parse_expression

type inToPost struct {
	stack  *operatorsStack
	inStr  string
	outStr string
}

func NewInToPost(inStr string) *inToPost {
	return &inToPost{
		stack:  NewOperandsStack(len(inStr)),
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
			this.stack.push(string(char))
		case string(')'):
			this.gotParen()
		default:
			this.outStr += string(char)
		}
	}
	for !this.stack.isEmpty() {
		this.outStr += this.stack.pop()
	}
}

func (this *inToPost) gotOp(op string, priority1 int) {
	for !this.stack.isEmpty() {
		stOp := this.stack.pop()
		if stOp == string('(') {
			this.stack.push(stOp)
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
				this.stack.push(stOp)
				break
			} else {
				this.outStr += stOp
			}
		}
	}
	this.stack.push(op)
}

func (this *inToPost) gotParen() {
	for !this.stack.isEmpty() {
		el := this.stack.pop()
		if el == string('(') {
			break
		} else {
			this.outStr += el
		}
	}
}
