package main

import "math"

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (c Cicle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (t Trapezium) Perimeter() float64 {
	return t.A + t.B + t.C + t.D
}