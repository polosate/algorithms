package exercises

func isUnique(input string) bool {
	var exists map[rune]bool

	for _, v := range []rune(input) {
		if exists[v] {
			return false
		}
		exists[v] = true
	}
	return true
}
