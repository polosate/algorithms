package huffman_code

import (
	"io/ioutil"
	"testing"
)

func TestNewHuffmanCode(t *testing.T) {
	text, err := ioutil.ReadFile("/home/ashaposhnikova/output/input")
	if err != nil {
		panic(err)
	}

	h := New()

	encoded := h.Encode(text)
	err = ioutil.WriteFile("/home/ashaposhnikova/output/out", encoded, 0666)
	if err != nil {
		panic(err)
	}
}
