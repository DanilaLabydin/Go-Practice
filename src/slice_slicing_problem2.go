package main

import (
	"fmt"
)

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	fmt.Println(y) // 1, 2, 30, 40, 50
	fmt.Println(x) // 1, 2, 30, 40
	fmt.Println(z) // 30, 40

	fmt.Println()
	x = append(x, 60)
	fmt.Println(y) // 1, 2, 30, 40, 60	
	fmt.Println(x) // 1, 2, 30, 40, 60
	fmt.Println(z) // 30, 40

	fmt.Println()
	z = append(z, 70)
	fmt.Println(y) // 1, 2, 30, 40, 70	
	fmt.Println(x) // 1, 2, 30, 40, 70
	fmt.Println(z) // 30, 40, 70

	secondSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	secondSlice2 := secondSlice[2: 5]
	fmt.Println(secondSlice)
	fmt.Println(secondSlice2)
	secondSlice2 = append(secondSlice2, 12345)
	fmt.Println(secondSlice)
	fmt.Println(secondSlice2)

}