package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
