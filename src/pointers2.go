package main

import (
	"fmt"
)

func main() {
	x := 10
	p := &x

	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)

	*p = 15
	fmt.Println(x)

	increment(&x)
	fmt.Println("Result: ", x)

	increment2(x)
	fmt.Println(x)

	var z *int
	fmt.Println(&z)
}

func increment(p *int) {
	*p += 1
}

func increment2(p int) {
	p += 1
}
