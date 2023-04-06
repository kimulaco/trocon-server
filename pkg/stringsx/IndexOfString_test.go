package stringsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOfString(t *testing.T) {
	assert.Equal(t, IndexOfString("2", []string{"1", "2", "3"}), 1)

	assert.Equal(t, IndexOfString("5", []string{"1", "2", "3"}), -1)
	assert.Equal(t, IndexOfString("2", []string{}), -1)
}

func BenchmarkIndexOfString(b *testing.B) {
	ss := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if index := IndexOfString("19", ss); index != 18 {
			b.Fatalf("Unexpected status IndexOfString: %d", index)
		}
	}
}
