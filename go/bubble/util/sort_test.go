package util_test

import (
	"swp/bubble/util"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	slice := util.GenerateSlice(20)
	util.Bubblesort(slice)
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
