package main
import "fmt"
import "unicode/utf8"

func main ()  {
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
	res2 := 3/2
	fmt.Println(res2) // 1

	// complex64, complex128
	var complexNum complex64 = 12 + 78i
	fmt.Println("complexNum", complexNum) // (12+78i)

	// string
	var name string = "John h"
	fmt.Println(name)
	fmt.Println(len("γ")) // 2
	fmt.Println(utf8.RuneCountInString("γ")) // 1

	// rune
	myRune := 'a'
	fmt.Println(myRune) // 97
}