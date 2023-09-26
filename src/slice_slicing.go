package main

import (
	"fmt"
)

func main() {
	var justSlice = []int{1, 2, 3, 4}
	y := justSlice[:2] // 1 2
	z := justSlice[1:] // 2 3 4
	d := justSlice[1:3] // 2 3 
	e := justSlice[:] // 1 2 3 4
	fmt.Println("Origin")
	fmt.Println("initial slice:", justSlice)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println()

	fmt.Println("change 'slice[1] = 20'")
	justSlice[1] = 20 // 1 20 3 4
	fmt.Println("initial slice:", justSlice) // 1 20 3 4
	fmt.Println("y:", y) // 1 20
	fmt.Println("z:", z) // 20 3 4
	fmt.Println("d:", d) // 20 3
	fmt.Println("e:", e) // 1 20 3 4
	fmt.Println()

	y[0] = 10
	fmt.Println("chahe 'y[0] = 10'")
	fmt.Println("initial slice:", justSlice) // 10 20 3 4
	fmt.Println("y:", y) // 10 20
	fmt.Println("z:", z) // 20 3 4
	fmt.Println("d:", d) // 20 3
	fmt.Println("e:", e) // 10 20 3 4
	fmt.Println()

	z[1] = 30
	fmt.Println("change 'z[1] = 30")
	fmt.Println("initial slice:", justSlice) // 10 20 30 4
	fmt.Println("y:", y) // 10 20
	fmt.Println("z:", z) // 20 30 4
	fmt.Println("d:", d) // 20 30
	fmt.Println("e:", e) // 10 20 30 4
	fmt.Println()

	x := []int{1, 2, 3, 4}
	y = x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, 123)
	fmt.Println(x, y)


	xx := make([]int, 4, 8)
	yy := xx[:2]
	fmt.Println(cap(xx), cap(yy))
	yy = append(yy, 123)
	fmt.Println(xx, yy)
	xx = append(xx, 321)
	fmt.Println(xx, yy)

}