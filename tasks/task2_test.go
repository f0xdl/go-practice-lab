package tasks

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"slices"
	"sort"
	"testing"
)

func TestProcessData(t *testing.T) {
	params := make([]int, 20000)
	for i := 0; i < len(params); i++ {
		params[i] = rand.Int()
	}
	processed := ProcessData(3, params)
	sortingParams := slices.Clone(params)
	for i := 0; i < len(sortingParams); i++ {
		sortingParams[i] *= sortingParams[i]
	}
	sort.Ints(sortingParams)
	for i := 0; i < len(params); i++ {
		assert.Equal(t, processed[i], sortingParams[i])
	}
}

func TestSortingSlice(t *testing.T) {
	params := make([]int, 10)
	for i := 0; i < len(params); i++ {
		params[i] = rand.Int()
	}
	//
	customParams := slices.Clone(params)
	SortingSlice(customParams)
	assert.NotEqual(t, params, customParams)
	//
	sortInt := slices.Clone(params)
	sort.Ints(sortInt)
	assert.NotEqual(t, params, sortInt)
	//
	assert.Equal(t, customParams, sortInt)
}
