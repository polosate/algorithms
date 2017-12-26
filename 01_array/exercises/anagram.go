package exercises

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	sArray := map[byte]int{}
	for _, v := range []byte(s1) {
		sArray[v]++
	}

	for _, v := range []byte(s2) {
		sArray[v]--
	}

	for _, v := range sArray {
		if v != 0 {
			return false
		}
	}
	return true
}
