package stringsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert.Equal(t, Contains("2", []string{"1", "2", "3"}), true)

	assert.Equal(t, Contains("5", []string{"1", "2", "3"}), false)
	assert.Equal(t, Contains("2", []string{}), false)
}

func BenchmarkContains(b *testing.B) {
	ss := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if isContains := Contains("19", ss); isContains != true {
			b.Fatalf("Unexpected status Contains: %t", isContains)
		}
	}
}
