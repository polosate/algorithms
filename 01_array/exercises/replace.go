package exercises

import "strings"

func replace(input string) (output string) {
	w := strings.NewReplacer(" ", "%20")
	return w.Replace(input)
}
