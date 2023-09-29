package main

import (
	"fmt"
)

func main() {
	test, test2 := "f", "d"
	fmt.Println(test, test2)

	for i := 1 ; i <= 100; i++ {
		var output string
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if len(output) != 0 {
			fmt.Println(output)
		} else {
			fmt.Println(i)
		}
	}
}