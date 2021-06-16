package main

import (
	"fmt"
	"swp/util"
	"time"
)

func main() {
	start := time.Now()
	bubbleTimes := []time.Duration{}
	insertionTimes := []time.Duration{}
	radixTimes := []time.Duration{}

	for i := 0; i < 101; i++ {
		fmt.Printf("\033[2K\r%d", i)
		slice := util.GenerateSlice(10000)
		startbubbleSort := time.Now()
		util.Bubblesort(slice)
		endbubbleSort := time.Now()

		bubbleTimes = append(bubbleTimes, endbubbleSort.Sub(startbubbleSort))

		startinsertionSort := time.Now()
		util.InsertionSort(slice)
		endinsertionSort := time.Now()
		insertionTimes = append(insertionTimes, endinsertionSort.Sub(startinsertionSort))

		radixSort := time.Now()
		util.RadixSort(slice)
		endradixSort := time.Now()
		radixTimes = append(radixTimes, endradixSort.Sub(radixSort))
	}
	fmt.Println()
	fmt.Println()
	end := time.Now()
	fmt.Println("Duration: ", end.Sub(start))
	fmt.Println("Array Length: ", 10000)

	fmt.Println("Bubblesort avg duration:", calcAvg(bubbleTimes))
	fmt.Println("Inseration avg duration:", calcAvg(insertionTimes))
	fmt.Println("Radix avg duration:", calcAvg(radixTimes))

}

func calcAvg(arr []time.Duration) time.Duration {
	sum := float64(0)
	for _, item := range arr {
		sum += float64(item.Nanoseconds())
	}
	avg := sum / float64(len(arr))
	return time.Duration(avg)

}
