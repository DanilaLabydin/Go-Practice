package main

import (
	"fmt"
)

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	evenVals2 := []int{2, 4, 6, 8, 10, 12}
	for i := range evenVals2 {
		fmt.Println(i)
	}

	evenVals3 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals3 {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k, v := range uniqueNames {
		fmt.Println(k, v)
	}

	for i := 0; i < 3; i ++ {
		fmt.Println()
	}

	for i := 1; i <= 3; i ++ {
		for k, v := range uniqueNames {
			fmt.Println(k, v)
		}
		fmt.Println()
	}

	simple_string := "Hello World"
	for i, r := range simple_string {
		fmt.Println(i, string(r))
	}
}