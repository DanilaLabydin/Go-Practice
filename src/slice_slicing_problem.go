package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 3, 3)
	slice2 := []int{0, 0, 0}
	slice3 := []int{}
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice1 = append(slice1, 1, 2)
	slice2 = append(slice2, 1, 2)
	slice3 = append(slice3, 1, 2)

	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3, len(slice3), cap(slice3))
}
