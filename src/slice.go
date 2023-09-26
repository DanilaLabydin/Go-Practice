package main

import (
	"fmt"
)

func main() {
	var simpleSlice []int
	fmt.Println(simpleSlice)
	var slice = []int{1, 2, 3, 4}
	fmt.Println(slice)
	fmt.Println(simpleSlice == nil)
	fmt.Println(slice == nil)
	fmt.Println("Get slices len:")
	fmt.Println(len(simpleSlice), len(slice))
	// add item to slice
	simpleSlice = append(simpleSlice, 12313)
	fmt.Println(simpleSlice)

	x := make([]int, 5)
	var y = make([]int, 5)
	fmt.Println(x)
	fmt.Println(y)

}
