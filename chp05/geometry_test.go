package geometry

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, expected float64) {
		got := shape.Perimeter()
		AssertGeometryFunctions(t, expected, got)
	}
	t.Run("Get perimeter from rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)
	})
	t.Run("Get perimeter from circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, math.Pi*2*10.0)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		got := shape.Area()
		AssertGeometryFunctions(t, expected, got)
	}
	t.Run("Get area from rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})
	t.Run("Get area from circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, math.Pi*10.0*10.0)
	})
}

func AssertGeometryFunctions(t testing.TB, expected, got float64) {
	t.Helper()
	if got != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, got)
	}
}
