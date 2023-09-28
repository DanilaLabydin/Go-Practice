package main

import (
	"fmt"
)

func main() {
	var str string = "Hello there"
	var letter byte = str[0]
	fmt.Println(str)
	fmt.Println(string(letter))
	var somePart string = str[4:7]
	fmt.Println(somePart)
	fmt.Println(len(str))

	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println()
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(s2)
	fmt.Println()
	var x int = 65
	var y = string(x)
	fmt.Println(y)
	y = fmt.Sprint(x)
	fmt.Println(y)

	var str_slice []byte = []byte(str)
	str_slice2 := []byte(str)
	fmt.Println(str_slice)
	fmt.Println(str_slice2)

}
