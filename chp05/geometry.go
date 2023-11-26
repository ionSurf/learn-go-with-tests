package geometry

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	// return 2 * (rectangle.Width + rectangle.Height)
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	// return rectangle.Width * rectangle.Height
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
