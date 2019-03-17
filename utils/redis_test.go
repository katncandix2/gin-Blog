package utils_test

import (
	"simpleBlog/utils"
	"testing"
)

func TestGetRedis(t *testing.T) {
	r := utils.GetRedis()
	if r == nil {
		panic("err")
	}
}

func BenchmarkGetRedis(b *testing.B) {

	for i := 0; i < b.N; i++ {
		utils.GetRedis()
	}
}
