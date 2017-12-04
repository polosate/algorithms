package huffman_code

import (
	"fmt"
	"testing"
)

func TestNewHuffmanCode(t *testing.T) {
	h := NewHuffmanCode("susi say that it is easy")
	h.makeFrequencyMap()
	fmt.Println(h.frequencyMap)
}
