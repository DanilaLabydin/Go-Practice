package main

import "fmt"

const pi = 3.14

func main() {
	var circleRadius int32 = 2
	circleArea := float32(circleRadius) * float32(circleRadius) * pi

	fmt.Println("The circle radius is:", circleRadius)
	fmt.Printf("The area is:", circleArea)

}
