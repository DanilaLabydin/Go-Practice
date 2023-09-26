package main

import (
	"fmt"
)

func main() {
	var array = [4]int{1, 2, 3, 4}
	slice1 := array[:2]
	slice2 := array[2:]
	fmt.Println(array, slice1, slice2)
	array[1] = 123
	fmt.Println(array, slice1, slice2)

}