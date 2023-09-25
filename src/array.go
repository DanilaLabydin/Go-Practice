package main

import (
	"fmt"
)

func main() {
	var todoList = [3]string{
		"buy bread",
		"buy milk",
		"buy bear",
	}

	fmt.Printf("The len(array) is: %d\n", len(todoList))

	for index, item := range todoList {
		fmt.Printf("%d. %s\n", index, item)
	}

	fmt.Println()
	var array1 = [3]string{}
	array1[0] = "1"
	array1[1] = "2"
	array1[2] = "3"

	for index, item := range array1 {
		fmt.Printf("%d. %s\n", index, item)
	}

	var array2 = [...]string{"aaaa"}
	fmt.Println(array2[0])

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	fmt.Println()
	i := 0
	for {
		if i == 15 {
			break
		}

		if i == 5 {
			break
		}

		fmt.Println(i)
		i++
	}

	var arr [30]int
	fmt.Println(arr)

}
