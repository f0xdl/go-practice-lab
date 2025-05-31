package tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounterCheck(t *testing.T) {
	result := CounterCheck(5)
	assert.Equal(t, true, result, "method return true")
}

func BenchmarkCounterCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CounterCheck(b.N)
	}
}
