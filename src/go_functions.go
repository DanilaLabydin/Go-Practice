package main

import "fmt"

const pi = 3.14

func main() {
	printCircleArea()
	fmt.Println("fdgdfgfdg")
}

func printCircleArea() {
	circleRadius := 2
	circleArea := float32(circleRadius) * float32(circleRadius) * pi

	fmt.Printf("The circle Radius: %d cm, \n", circleRadius)
	fmt.Println("The formula to compute the circle area: S = PR2\n")
	fmt.Printf("The circle area: %f32\n", circleArea)
	fmt.Printf("The circle area: %f32\n", circleArea)
}