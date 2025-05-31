package tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUniqueValues(t *testing.T) {
	input := []int{1, 2, 2, 1, 4, 3}
	output := []int{1, 2, 4, 3}

	uniqueValues := GetUniqueValues(input)

	for i, v := range uniqueValues {
		if v != output[i] {
			assert.Equal(t, output[i], v)
		}
	}
}

func TestGetUniqueValuesMap(t *testing.T) {
	input := []int{1, 2, 2, 1, 4, 3}
	output := []int{1, 2, 4, 3}

	uniqueValues := GetUniqueValuesMap(input)

	for i, v := range uniqueValues {
		assert.Equal(t, output[i], v)
	}
}
