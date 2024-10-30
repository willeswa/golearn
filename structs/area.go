package main

import "math"

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Cicle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Trapezium) Area() float64 {
	return 0.0
}