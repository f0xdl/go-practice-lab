package tasks

func CharFrequency(s string) map[rune]int {
	result := make(map[rune]int)
	for _, v := range s {
		result[v]++
	}
	return result
}
