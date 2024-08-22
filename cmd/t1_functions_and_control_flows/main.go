package main

import (
	"errors"
	"fmt"
)

func main() {
	x := 8
	y := 3
	z := add(x, y) // 11
	println(z)

	res, rem, err := divide(x, y)
	// &&, ||, !, ==, !=, <, <=, >, >= are the same as in other languages
	// if, else if, else
	if err != nil {
		fmt.Println(err)
	} else if rem == 0 {
		fmt.Println("rem is zero")
	} else {
		fmt.Printf("res: %d, rem: %d\n", res, rem) // eg. res: 2, rem: 2
		fmt.Printf("res: %v, rem: %v\n", res, rem) // eg. res: 2, rem: 2
	}

	// switch, case, default // break is not needed
	switch {
	case err != nil:
		fmt.Println(err)
	case rem == 0:
		fmt.Println("rem is zero")
	default:
		fmt.Printf("res: %d, rem: %d\n", res, rem) // res: 2, rem: 2
	}

	// conditional switch
	switch rem {
	case 0:
		fmt.Println("rem is zero")
	case 1:
		fmt.Println("rem is one")
	default:
		fmt.Println("rem is not zero or one, it is: ", rem)
	}

}

func add(x, y int) int {
	return x + y
}
func divide(x, y int) (int, int, error) {
	var err error // has default value of nil
	if y == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	// return 0, 0
	return x / y, x % y, err
}
