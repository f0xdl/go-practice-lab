package tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharFrequency(t *testing.T) {
	input := "Hello, 世界"
	expected := map[rune]int{
		'H': 1, 'e': 1, 'l': 2, 'o': 1, ',': 1, ' ': 1,
		'世': 1, '界': 1,
	}
	assert.Equal(t, expected, CharFrequency(input))
}
