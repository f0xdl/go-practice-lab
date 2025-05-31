package tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSlice(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	o := []int{1, 2, 3, 4, 5, 6}

	result := MergeSlice(a, b)
	assert.Equal(t, result, o)
}

func TestMergeSliceALen(t *testing.T) {
	a := []int{1, 3, 3, 5}
	b := []int{2, 4, 6}
	o := []int{1, 2, 3, 3, 4, 5, 6}

	result := MergeSlice(a, b)
	assert.Equal(t, result, o)
}

func TestMergeSliceBLen(t *testing.T) {
	a := []int{1, 3, 3, 5}
	b := []int{2, 4, 6, 7, 8, 9, 9}
	o := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 9}

	result := MergeSlice(a, b)
	assert.Equal(t, result, o)
}

func TestMergeSliceSimplify(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	o := []int{1, 2, 3, 4, 5, 6}

	result := MergeSliceSimplify(a, b)
	assert.Equal(t, result, o)
}

func TestMergeSliceSimplifyALen(t *testing.T) {
	a := []int{1, 3, 3, 5}
	b := []int{2, 4, 6}
	o := []int{1, 2, 3, 3, 4, 5, 6}

	result := MergeSliceSimplify(a, b)
	assert.Equal(t, result, o)
}

func TestMergeSliceSimplifyBLen(t *testing.T) {
	a := []int{1, 3, 3, 5}
	b := []int{2, 4, 6, 7, 8, 9, 9}
	o := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 9}

	result := MergeSliceSimplify(a, b)
	assert.Equal(t, result, o)
}

func TestMergeSliceSimplifyWithNegatives(t *testing.T) {
	a := []int{-10, -3, 0, 2}
	b := []int{-5, -1, 1, 3}
	expected := []int{-10, -5, -3, -1, 0, 1, 2, 3}

	result := MergeSliceSimplify(a, b)

	assert.Equal(t, expected, result)
}
