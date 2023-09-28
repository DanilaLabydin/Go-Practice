package main

import (
	"fmt"
)

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	bob := person{}
	fmt.Println(fred)
	fmt.Println(bob)

	julia := person{
		"Julia",
		40,
		"cat",
	}

	beth := person{
		age:  30,
		name: "Betch",
	}
	fmt.Println(julia)
	fmt.Println(beth)
}
