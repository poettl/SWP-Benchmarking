package main

import (
	"fmt"
	"swp/insertion/util"
)

func main() {
	slice := util.GenerateSlice(20)
	fmt.Println("\n Unsorted \n\n", slice)
	util.InsertionSort(slice)
	fmt.Println("\n Sorted \n\n", slice)
}
