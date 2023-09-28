package main

import (
	"fmt"
)

func main() {
	var x [3]int
	fmt.Println(x) // [0, 0, 0]

	var y = [3]int{20, 30, 40}
	fmt.Println(y)

	var z = [12]int{1, 5: 4, 6, 10: 100, 15} // [1 0 0 0 0 4 6 0 0 0 100 15]
	fmt.Println(z)
}
