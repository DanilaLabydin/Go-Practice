package main

import "fmt"

const pi = 3.14

func main() {
	var circleRadius int = 2
	var circleArea float32 = caclulateCircleArea(circleRadius)
	printResults(circleRadius, circleArea)
}

func caclulateCircleArea(radius int) float32 {
	return float32(radius) * float32(radius) * pi
}

func printResults(radius int, area float32) {
	fmt.Printf("The circle radius: %d\n", radius)
	fmt.Printf("The circle area: %f32\n", area)
}