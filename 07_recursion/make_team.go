package _7_recursion

func makeTeams(prefix string, group string, size int) []string {
	groupB := []byte(group)
	if size == 0 {
		return []string{prefix}
	}
	if len(groupB) == 0 {
		return []string{}
	}
	result := []string{}
	result = append(result, makeTeams(prefix+string(groupB[0]), string(groupB[1:]), size-1)...)
	result = append(result, makeTeams(prefix, string(groupB[1:]), size)...)
	return result
}
