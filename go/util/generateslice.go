package util

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) []int32 {

	slice := make([]int32, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		v1 := rand.Intn(999)
		v2 := rand.Intn(999)
		slice[i] = int32(v1 - v2)
	}
	return slice
}
