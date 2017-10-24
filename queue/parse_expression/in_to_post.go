package parse_expression

type inToPost struct {
	stack *operatorsStack
	inBytes   []byte
	outStr  string
}

func NewInToPost(inStr string) *inToPost {
	inputBytes := []byte(inStr)
	return &inToPost{
		stack:        NewOperandsStack(len(inputBytes)),
		inBytes:   inputBytes,
		outStr: "",
	}
}

func (this *inToPost) DoTrans() {
	
}
