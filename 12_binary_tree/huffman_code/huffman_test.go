package huffman_code

import (
	"fmt"
	"testing"
	//"unsafe"
)

func TestNewHuffmanCode(t *testing.T) {
	text := "susi say that it is easy"
	res := ""
	for _, c := range text {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	fmt.Println("===src===")
	fmt.Println(text)
	fmt.Println(res)
	fmt.Println("===src===")

	h := New()

	encoded := h.Encode(text)
	fmt.Println("===encoded===")
	fmt.Println(encoded)
	fmt.Println("===encoded===")

	fmt.Println("===decoded===")
	fmt.Println(h.decode(encoded))
	fmt.Println("===decoded===")

	//var x1 []uint8
	//fmt.Printf("x1: %T, %d\n", x1, unsafe.Sizeof(x1))
	//var x2 uint8
	//fmt.Printf("x2: %T, %d\n", x2, unsafe.Sizeof(x2))
	//var x3 uint16
	//fmt.Printf("x3: %T, %d\n", x3, unsafe.Sizeof(x3))
	//var x4 uint32
	//fmt.Printf("x4: %T, %d\n", x4, unsafe.Sizeof(x4))
	//var x5 uint64
	//fmt.Printf("x5: %T, %d\n", x5, unsafe.Sizeof(x5))
	//var b byte
	//fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	//var b2 []byte
	//fmt.Printf("b2: %T, %d\n", b2, unsafe.Sizeof(b2))
	//var s string
	//fmt.Printf("s: %T, %d\n", s, unsafe.Sizeof(s))
	//var s1 []string
	//fmt.Printf("s1: %T, %d\n", s1, unsafe.Sizeof(s1))
	//
	//x2 = 255
	//fmt.Println(x2)

}
