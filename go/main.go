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
	end := time.Now()
	sumB, avgB := calcAvg(bubbleTimes)
	sumI, avgI := calcAvg(insertionTimes)
	sumR, avgR := calcAvg(radixTimes)

	fmt.Println("Duration: ", end.Sub(start))
	fmt.Println("Array Length: ", 10000)
	fmt.Println()
	fmt.Println("--- Bubble Sort ---")
	fmt.Println("Total: ", sumB)
	fmt.Println("Avg: ", avgB)
	fmt.Println()
	fmt.Println("--- Inseration Sort ---")
	fmt.Println("Total: ", sumI)
	fmt.Println("Avg: ", avgI)
	fmt.Println()
	fmt.Println("--- Radix Sort ---")
	fmt.Println("Total: ", sumR)
	fmt.Println("Avg: ", avgR)

}

func calcAvg(arr []time.Duration) (time.Duration, time.Duration) {
	sum := float64(0)
	for _, item := range arr {
		sum += float64(item.Nanoseconds())
	}
	avg := sum / float64(len(arr))
	return time.Duration(sum), time.Duration(avg)

}
