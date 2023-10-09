package main

import (
	"fmt"
	"math"
)

/* Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.*/

type Point struct {
	x int
	y int
}

func newPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func squareCalculate(a Point, b Point) float64 {
	dx := math.Abs(float64(a.x - b.x))
	dy := math.Abs(float64(a.y - b.y))
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func main() {
	a := newPoint(1, 2)
	b := newPoint(10, 20)
	fmt.Println(squareCalculate(a, b))
}
