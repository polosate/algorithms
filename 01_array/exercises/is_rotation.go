package exercises

import "strings"

func isRotation(s1, s2 string) bool {
	if len(s1) != len(s2) || len(s1) == 0 {
		return false
	}
	return strings.Contains(s1+s1, s2)
}
