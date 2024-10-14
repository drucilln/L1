package main

import (
	"fmt"
	"math"
)

type Point struct { // Внутри одного пакета, даже приватные поля будут доступны
	x, y float64
}

func Distance(p1, p2 Point) float64 { // AB = √(xb - xa)2 + (yb - ya)2
	ac := p1.x - p2.x              // AC = xb - xa;
	bc := p1.y - p2.y              //BC = yb - ya.
	ab := math.Sqrt(ac*ac + bc*bc) //AB = √AC2 + BC2.

	return ab
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func main() {
	point1 := NewPoint(2, 3)
	point2 := NewPoint(3, 4)

	dist := Distance(point1, point2)
	fmt.Printf("%.4f\n", dist)
}
