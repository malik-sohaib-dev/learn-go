package main

import (
	"fmt"
	"time"
)

func main() {
	var myMap = map[string]int32{"one": 1, "two": 2, "three": 3, "four": 4}
	for key, value := range myMap {
		fmt.Printf("Key: %v, Value: %v | ", key, value)
	}
	fmt.Println()
	var mySlice = []int32{1, 2, 3, 4}
	for index, value := range mySlice {
		fmt.Printf("Index: %v, Value: %v | ", index, value)
	}
	fmt.Println()
	var myArray = [4]int32{1, 2, 3, 4}
	for index, value := range myArray {
		fmt.Printf("Index: %v, Value: %v | ", index, value)
	}
	fmt.Println()
	for i := 0; i < 10; i++ { // for init; condition; post
		if i%2 == 0 {
			continue
		} else if i == 7 {
			break
		}

		fmt.Printf("%v ", i)
	}
	fmt.Println()
	// Slice append cost analysis
	const size = 100000000
	var slice1 = []int{}
	var slice2 = make([]int, 0, size)
	fmt.Printf("Slice1 append cost: %v\n", sliceAppendCost(slice1, size)) // 15.2377ms
	fmt.Printf("Slice2 append cost: %v\n", sliceAppendCost(slice2, size)) // 4.3468ms
}

func sliceAppendCost(slice []int, size int) time.Duration {
	ts := time.Now()
	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}
	fmt.Println(len(slice))
	return time.Since(ts)
}
