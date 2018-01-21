package exercises

func twoSum(nums []int, target int) []int {
	newMap := map[int]int{}
	for i, v := range nums {
		complement := target - v
		if index, ok := newMap[complement]; ok {
			return []int{i, index}
		}
		newMap[v] = i
	}
	return []int{}
}
