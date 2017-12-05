package huffman_code

import (
	"strings"
)

type huffmanCode struct {
	text         string
	frequencyMap map[string]int
	codeMap      map[string]byte
}

func NewHuffmanCode(text string) huffmanCode {
	return huffmanCode{
		text:         text,
		frequencyMap: map[string]int{},
	}
}

func (h *huffmanCode) makeFrequencyMap() {
	byteArr := []byte(h.text)
	var currentChar string
	for i := range byteArr {
		currentChar = string(byteArr[i])
		if h.frequencyMap[currentChar] > 0 {
			continue
		} else {
			h.frequencyMap[currentChar] = strings.Count(h.text, currentChar)
		}
	}
}
