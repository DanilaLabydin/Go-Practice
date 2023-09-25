package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(z, d)

	var m, n int = 20, 25
	fmt.Println(m, n)
}
