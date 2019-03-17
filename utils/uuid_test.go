package utils_test

import (
	"simpleBlog/utils"
	"testing"
)

func TestUUID(t *testing.T) {
	utils.GetUUID()
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.GetUUID()
	}
}
