package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// int, int8, int16, int32, int64 (Its the bit size of the int)
	var intNum int8 = 12 + 78
	fmt.Println(intNum)

	// float32, float64
	var floatNum float32 = 12.78
	fmt.Println(floatNum)
	fmt.Println("Hello, World!")

	// type casting
	res := floatNum + float32(intNum)
	fmt.Println(res) // 90.78
	res2 := 3 / 2
	fmt.Println(res2) // 1

	// complex64, complex128
	var complexNum complex64 = 12 + 78i
	fmt.Println("complexNum", complexNum) // (12+78i)

	// string
	var name string = "John h"
	fmt.Println(name)
	fmt.Println(len("γ"))                    // 2
	fmt.Println(utf8.RuneCountInString("γ")) // 1 - utf8.RuneCountInString counts the number of runes in a string

	// rune - alias for int32 - kind of like a character
	myRune := 'a'
	fmt.Println(myRune) // 97

	// Advanced Strings
	var myString = "Résumé" // é is a non-ASCII character
	var indexed = myString[0]

	fmt.Printf("%v, %T\n", indexed, indexed) // 82, uint8
	fmt.Println(len(myString))               // 8
	for i, j := range myString {
		// range converts the string to a rune slice and returns the index and the rune, thus the skipping of indexes for non-ASCII characters
		fmt.Println(i, j)
	}
	// String encoding history
	// ASCII - 7 bits - 128 characters - But ASCII is only for English, not special characters
	// UTF-32 - 32 bits - 4 bytes - Fixed length - Waste of space
	// UTF-8 - 8 bits - 1 byte - Variable length - ASCII is a subset of UTF-8

	// type casting from string to rune
	myRune2 := []rune(myString)
	for i, j := range myRune2 {
		// In this case, we'll get the correct indexes for non-ASCII characters
		fmt.Println(i, j)
	}

	// String building
	var strSlice = []string{"Hello", "World"}
	var concatinated = ""
	for _, v := range strSlice {
		// This is inefficient as strings are immutable and a new string is created everytime a string is concatinated
		concatinated += v
	}
	fmt.Println(concatinated) // HelloWorld
	// concatinated[1] = 'a' // This will throw an error
	// String can't be mutated, but can be converted to a slice of runes and then mutated

	// Efficient string building
	var strSliceEffieint = []string{"Hello", "World"}
	var stringBuilder strings.Builder
	for _, v := range strSliceEffieint {
		// Array is allocated internally and the string is concatinated to the array
		stringBuilder.WriteString(v)
	}
	var concatinatedEfficient = stringBuilder.String()
	fmt.Println(concatinatedEfficient) // HelloWorld
}
