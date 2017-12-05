package huffman_code

import (
	"fmt"
	"testing"
)

func TestNewHuffmanCode(t *testing.T) {
	h := NewHuffmanCode("susi say that it is easy\n stop it\n")
	h.makeFrequencyMap()
	fmt.Println(h.frequencyMap)
}
