package tasks

import "math"

func MergeSlice(a, b []int) []int {
	lenAB := len(a) + len(b)
	result := make([]int, lenAB)
	iA, iB := 0, 0
	for i := 0; i < lenAB; i++ {
		tempA := math.MaxInt
		if iA < len(a) {
			tempA = a[iA]
		}
		tempB := math.MinInt
		if iB < len(b) {
			tempB = b[iB]
		}
		if tempA > tempB {
			result[i] = tempB
			iB++
		} else if tempA < tempB {
			result[i] = tempA
			iA++
		} else {
			result[i] = tempA
			iA++
			i++
			result[i] = tempB
			iB++
		}
	}
	return result
}

func MergeSliceSimplify(a, b []int) []int {
	var result []int
	iA, iB := 0, 0

	for iA < len(a) && iB < len(b) {
		if a[iA] < b[iB] {
			result = append(result, a[iA])
			iA++
		} else {
			result = append(result, b[iB])
			iB++
		}
	}
	if iA < len(a) {
		result = append(result, a[iA:]...)
	}
	if iB < len(b) {
		result = append(result, b[iB:]...)
	}
	return result
}
