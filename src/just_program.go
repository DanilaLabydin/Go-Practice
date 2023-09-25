package main

import "fmt"

var globalVar string = "I am a global var"

func main() {
	var globalVar int = 2123213
	local_var := 232
	local_var = 2

	var loc_var string = "I am a local var"
	fmt.Println(local_var)
	fmt.Println(globalVar)
	fmt.Println(loc_var)
	fmt.Println(globalVar)
}
