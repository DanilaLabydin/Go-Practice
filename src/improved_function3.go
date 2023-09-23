package main

import (
	"fmt"
	"errors"
)

const pi = 3.14

func main() {
	printCircleArea(-2)
}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Radius: %d\n", radius)
	fmt.Printf("Area: %f32\n", circleArea)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным!")
	}

	return float32(radius) * float32(radius) * pi, nil
}