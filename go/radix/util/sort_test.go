package util_test

import (
	"swp/radix/util"
	"testing"
)

func TestRadix(t *testing.T) {
	slice := util.GenerateSlice(20)
	util.RadixSort(slice)
	if !checkSortedArray(slice) {
		t.Errorf("Failed")
	}

}

func checkSortedArray(arr []int32) bool {
	sortedArray := true
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				sortedArray = false
				break
			}
		}
	}
	return sortedArray
}
