package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(intSlice)) // 15

	var floatSlice = []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumSlice(floatSlice)) // 16.5

	fmt.Println(isEmpty([]int{1, 2, 4}))     // false
	fmt.Println(isEmpty([]float64{}))        // true
	fmt.Println(isEmpty([]string{"a", "b"})) // false

	// JSON
	var contacts = loadJSON[contactInfo]("./cmd/t9_generics/contacts.json")
	fmt.Printf("%+v\n", contacts)

	var purchases = loadJSON[purchaseInfo]("./cmd/t9_generics/purchases.json")
	fmt.Println(purchases)
}

// Generics in Go: Generics are a way to write functions or data structures that can work with any data type

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Any type: any is a type that can hold any value
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ := os.ReadFile(filePath)
	fmt.Println(string(data))

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}
