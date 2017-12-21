package _5_hash_table

func groupSize(hashTableSize int) int {
	groupSize := 0
	for {
		if hashTableSize/10 < 1 {
			return groupSize
		}
		if hashTableSize%10 > 0 {
			groupSize++
		}
		groupSize++
		hashTableSize /= 10
	}
	return groupSize
}

func fold(key int64, hashTableSize int) int {
	return 0
}
