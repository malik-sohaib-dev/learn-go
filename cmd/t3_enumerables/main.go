package main

import (
	"fmt"
)

func main() {
	// arrays have fixed length, fixed value type, indexable and contiguous in memory
	var arr [4]int32 = [4]int32{1, 2, 3, 4}
	// var arr = [4]int32{1, 2, 3, 4}
	// arr := [4]int32{1, 2, 3, 4}
	// arr := [...]int32{1, 2, 3, 4}
	fmt.Println(arr)      // [1 2 3 4]
	fmt.Println(arr[0])   // 1
	fmt.Println(arr[1:4]) // [2 3 4] // [inclusive:exclusive]
	fmt.Println(arr[1:])  // [2 3 4]
	fmt.Println(arr[:3])  // [1 2 3]
	fmt.Println(len(arr)) // 4

	// slices - wrapper around arrays - dynamic size arrays - reference type
	var slice []int32 = []int32{1, 2, 3, 4}
	// var slice = []int32{1, 2, 3, 4}
	// slice := []int32{1, 2, 3, 4}
	fmt.Printf("Length is: %v, Capacity is: %v\n", len(slice), cap(slice))              // Length is: 4, Capacity is: 4
	slice = append(slice, 5)                                                            // append
	fmt.Printf("After append Length is: %v, Capacity is: %v\n", len(slice), cap(slice)) // Length is: 5, Capacity is: 8
	fmt.Println(slice)                                                                  // [1 2 3 4 5]
	fmt.Println(slice[0])                                                               // 1
	fmt.Println(slice[1:4])                                                             // [2 3 4] // [inclusive:exclusive]
	fmt.Println(slice[1:])                                                              // [2 3 4 5]
	fmt.Println(slice[:3])                                                              // [1 2 3]
	fmt.Println(len(slice))                                                             // 5
	slice = append(slice, []int32{1, 2, 3}...)                                          // unpacking - spread operator
	// slice = append(slice, 1,2,3) // same as above
	fmt.Println(slice) // [1 2 4 5]

	var makeSlice = make([]int32, 5, 10)                                           // make(type, length, capacity)
	fmt.Printf("Length is: %v, Capacity is: %v\n", len(makeSlice), cap(makeSlice)) // Length is: 5, Capacity is: 10

	// maps - key value pairs - reference type
	var m map[string]int32 = map[string]int32{"one": 1, "two": 2, "three": 3, "four": 4}
	fmt.Println(m) // map[four:4 one:1 three:3 two:2]
	var m2 map[string]int32 = make(map[string]int32)
	m2["one"] = 1
	fmt.Println(m2["one"]) // 1
	// if key is not present, it will default value of that type, 0 for int32
	fmt.Println(m2["two"]) // 0
	value, ok := m2["two"] // ok is false if key is not present
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key not present")
	}
	delete(m, "one") // delete key by reference so no value returned
	fmt.Println(m)   // map[two:0]
}
