package main

import (
	"fmt"
)

const x int64 = 10

const (
	idkey = "id"
	nameKey = "name"
)

const z = 20 * 20

func main() {
	var y int = int(x)


	fmt.Println(x)
	fmt.Println(y)

	y = 32424


	fmt.Println(x)
	fmt.Println(y)
}