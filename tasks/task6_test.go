package tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCountWords(t *testing.T) {
	input := "Hello, world!\nThis is Go."
	result := GetCountWords(input)
	assert.Equal(t, 5, result)

}

func TestGetCountWordsNewLine(t *testing.T) {
	input := "Hello, world!\nThis is Go.\n"
	result := GetCountWords(input)
	assert.Equal(t, 5, result)

}

func TestMultiSplit(t *testing.T) {
	input := "Hello, world!\nThis is Go."
	result := MultiSplit(input, " ", "i")
	assert.Equal(t, 6, len(result))
}

func TestMultiSplitEmpty(t *testing.T) {
	input := "Hello, world!\nThis is Go."
	result := MultiSplit(input, "")
	assert.Equal(t, 1, len(result))
}

func TestMultiSplitUpdated(t *testing.T) {
	input := "Hello, world!\nThis is Go."
	result := MultiSplitUpdated(input, " ", "i")
	assert.Equal(t, 6, len(result))
}

func TestMultiSplitUpdatedPanic(t *testing.T) {
	input := "Hello, world!\nThis is Go."
	assert.Panics(t, func() {
		MultiSplitUpdated(input, "")
	})
}

func TestMultiSplitUpdatedEmpty(t *testing.T) {
	input := "Hello, world!\nThis is Go."
	result := MultiSplitUpdated(input)
	assert.Equal(t, 1, len(result))
}
