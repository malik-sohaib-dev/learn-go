package main

import "fmt"

// Structs are a way to define a custom data type
type person struct {
	name      string
	age       uint8
	hobbyInfo hobby // This can be accessed using p1.hobbyInfo.name
	// hobby // This can be accessed directly using p1.name
	int
}

type hobby struct {
	name string
}

type animal struct {
	name string
	age  uint8
}

// Struct methods
func (p person) getName() string {
	return p.name
}

func (a animal) getName() string {
	return a.name
}

// Interface
type general interface {
	getName() string
}

func generalName(g general) string {
	return g.getName()
}

func main() {
	// var p1 person = person{name: "John", age: 23}
	var p1 person = person{"John", 23, hobby{"Haha"}, 10}
	p1.age = 25
	p1.hobbyInfo.name = "Hehe"
	fmt.Println(p1) // {John 25 {Hehe} 10}

	// Anonymous struct - not reuseable
	var p2 = struct {
		name string
		age  uint8
	}{"Broj", 23}
	fmt.Println(p2) // {Broj 23}

	var a1 = animal{"Bubbles", 5}
	fmt.Println("Name:", generalName(a1))
	fmt.Println("Name:", generalName(p1))
}
