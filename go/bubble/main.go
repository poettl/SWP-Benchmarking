package main

import (
	"fmt"
	"swp/bubble/util"
)

func main() {
	slice := util.GenerateSlice(20)
	fmt.Println("\n Unsorted \n\n", slice)
	util.Bubblesort(slice)
	fmt.Println("\n Sorted \n\n", slice)
}
