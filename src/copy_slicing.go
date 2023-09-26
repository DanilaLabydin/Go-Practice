package main

import (
	"fmt"
)

func main() {
	var initialSlice = []int{1, 2, 3, 4}
	var copySlice = make([]int, 4)
	copyQuantity := copy(copySlice, initialSlice)
	fmt.Println(initialSlice)
	fmt.Println(copySlice)
	fmt.Println(copyQuantity)

}
