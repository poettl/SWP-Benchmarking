package util

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) []int {

	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		v1 := rand.Intn(999)
		v2 := rand.Intn(999)
		slice[i] = v1 - v2
	}
	return slice
}
