package parse_expression

import "fmt"

type inToPost struct {
	stack   *operatorsStack
	inBytes []byte
	outStr  string
}

func NewInToPost(inStr string) *inToPost {
	inputBytes := []byte(inStr)
	return &inToPost{
		stack:   NewOperandsStack(len(inputBytes)),
		inBytes: inputBytes,
		outStr:  "",
	}
}

func (this *inToPost) DoTrans() {
	var strCh string
	for _, v := range this.inBytes {
		strCh = string(v)
		switch strCh {
		case "+", "-":
			fmt.Println()
		case "*", "/":
			fmt.Println()
		case "(":
			fmt.Println()
		case ")":
			fmt.Println()
		default:
			this.outStr += strCh
		}
	}
}
