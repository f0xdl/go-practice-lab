package tasks

func GetUniqueValues(input []int) []int {
	var result []int
	for _, v := range input {
		find := false
		for _, rv := range result {
			if v == rv {
				find = true
				break
			}
		}
		if !find {
			result = append(result, v)
		}
	}
	return result
}

func GetUniqueValuesMap(input []int) []int {
	result := make([]int, 0)
	exists := make(map[int]struct{})

	for _, v := range input {
		if _, found := exists[v]; !found {
			result = append(result, v)
			exists[v] = struct{}{}
		}
	}
	return result
}
