package geometry

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				Width:  10.0,
				Height: 10.0,
			},
			expected: 40.0,
		},
		{
			name: "Circle",
			shape: Circle{
				Radius: 10.0,
			},
			expected: math.Pi * 10.0 * 2,
		},
		{
			name: "Triangle",
			shape: Triangle{
				Base:   12,
				Height: 6,
			},
			expected: 0,
		},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			AssertGeometryFunctions(t, tt.shape, tt.expected, tt.shape.Perimeter())
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				Width:  12,
				Height: 6,
			},
			expected: 72.0,
		},
		{
			name: "Circle",
			shape: Circle{
				Radius: 10.0,
			},
			expected: math.Pi * 10.0 * 10.0,
		},
		{
			name: "Triangle",
			shape: Triangle{
				Base:   12,
				Height: 6,
			},
			expected: 36.0,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			AssertGeometryFunctions(t, tt.shape, tt.expected, tt.shape.Area())
		})
	}
}

func AssertGeometryFunctions(t testing.TB, shape Shape, expected, got float64) {
	t.Helper()
	if got != expected {
		t.Errorf("%v, expected %.2f, got %.2f", shape, expected, got)
	}
}
