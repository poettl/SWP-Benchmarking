package util_test

import (
	"swp/insertion/util"
	"testing"
)

func TestInsertion(t *testing.T) {
	slice := util.GenerateSlice(20)
	util.InsertionSort(slice)
	isSorted := checkSortedArray(slice)

	if !isSorted {
		t.Errorf("Failed")
	}

}

func checkSortedArray(arr []int) bool {
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
