package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Println(&x)
	if x > 5 {
		fmt.Println(&x)
		fmt.Println(x)
		x := 5
		fmt.Println(&x)
		fmt.Println(x)
	}
	fmt.Println(x)
	
}