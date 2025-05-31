package tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptySlice(t *testing.T) {
	input := []string{}
	result := CountStringLength(input)
	assert.Equal(t, 0, len(result))
}

func TestCountStringLength(t *testing.T) {
	input := []string{"go", "rust", "js", "go", "python", "ts"}
	result := CountStringLength(input)
	assert.Equal(t, 4, result[2])
	assert.Equal(t, 1, result[4])
	assert.Equal(t, 1, result[6])
}
