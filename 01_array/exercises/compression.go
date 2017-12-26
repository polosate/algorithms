package exercises

import "strconv"

// For ASCII
func compressString(input string) (output string) {
	byteInput := []byte(input)
	last := byteInput[0]
	count := 1
	for i := 1; i < len(byteInput); i++ {
		if byteInput[i] == last {
			count++
		} else {
			output += string(last) + strconv.Itoa(count)
			last = byteInput[i]
			count = 1
		}
	}
	output += string(last) + strconv.Itoa(count)
	if len(input) <= len(output) {
		return input
	} else {
		return output
	}
}
