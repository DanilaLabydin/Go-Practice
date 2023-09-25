package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)
	x := slice[0]
	slice[3] = x
	fmt.Println(slice)

}