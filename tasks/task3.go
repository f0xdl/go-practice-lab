package tasks

// CountStringLength map [len,count]
func CountStringLength(input []string) map[int]int {
	counts := map[int]int{}
	for _, v := range input {
		vLen := len(v)
		counts[vLen]++
	}
	return counts
}
