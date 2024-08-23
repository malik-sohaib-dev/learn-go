package main

import "fmt"

func main() {
	var p *int32              // stored nill, doesn't point to any memory address
	var q *int32 = new(int32) // stored the memory address of the int32 value
	// *q will give the value stored at the memory address which is 0 at this point
	// *p will give an error because it doesn't point to any memory address
	*q = 42 // set the value at the memory address to 42

	// *p = 42 // This will give an error because p doesn't point to any memory address

	var i int32 = 32
	q = &i // set the memory address of i to q

	fmt.Println(p, *q) // <nil> 32

	*q = 11
	fmt.Println(i) // 11

	var a int32 = 22
	fmt.Printf("Memory address of a: %v\n", &a)
	passbyValue(a)
	fmt.Println(a) // 22
	passbyReference(&a)
	fmt.Println(a) // 0
}

func passbyValue(a int32) {
	fmt.Printf("Pass by value memory address: %v\n", &a)
	a = 0
}

func passbyReference(a *int32) {
	fmt.Printf("Pass by reference memory address: %v\n", a)
	*a = 0
}
