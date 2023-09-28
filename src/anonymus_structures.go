package main

import (
	"func"
)

func main() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
}
