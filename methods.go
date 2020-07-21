package main

import (
	"fmt"
	"math"
)

type Vertex_ struct {
	X, Y float64
}

func (v Vertex_) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func Abs(v Vertex_) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	fmt.Println("Hello", "World")
	v := Vertex_{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
