package util_test

import (
	"swp/util"
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

func TestInsertion(t *testing.T) {
	slice := util.GenerateSlice(20)
	util.InsertionSort(slice)
	isSorted := checkSortedArray(slice)

	if !isSorted {
		t.Errorf("Failed")
	}

}

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
