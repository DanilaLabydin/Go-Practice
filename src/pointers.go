package main

import (
	"fmt"
)

var x_global int = 10

func main() {
	x := 10

	fmt.Println(x)
	fmt.Println(&x)

	fmt.Println(x_global)
	fmt.Println(&x_global)

	p := &x
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println()
	fmt.Println()

	x = 10
	p = &x
	fmt.Println("x:", x)
	fmt.Println("x address:", p)
	fmt.Println("p value:", *p)

	*p = 15
	fmt.Println("x:", x)

}
