package parse_expression

import (
	"testing"
	"fmt"
)

func TestParseCase01(t *testing.T) {
	pp :=  NewParsePost("1243-*+")
	res := pp.DoCalculate()
	fmt.Println(res)
}
