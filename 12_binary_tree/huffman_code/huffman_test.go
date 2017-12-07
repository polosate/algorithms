package huffman_code

import (
	"io/ioutil"
	"testing"
)

func TestNewHuffmanCode(t *testing.T) {
	path := "/home/ashaposhnikova/output/input2"
	text, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	h := New()

	encoded := h.Encode(text)
	err = ioutil.WriteFile(path+"_out", encoded, 0666)
	if err != nil {
		panic(err)
	}
}
