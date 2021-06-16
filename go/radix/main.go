package main

import (
	"fmt"
	"swp/radix/util"
)

func main() {
	slice := util.GenerateSlice(20)
	fmt.Println("\n Unsorted \n\n", slice)
	util.RadixSort(slice)
	fmt.Println("\n Sorted \n\n", slice)
}
