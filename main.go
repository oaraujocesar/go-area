package goarea

import "math"

// Pi is a numeric proportion defined by the relation between the perimeter of a circle
// and its diameter
const Pi = 3.1415

// Circle calculates the area of a circle
func Circle(ray float64) float64 {
	return math.Pow(ray, 2) * Pi
}

// Rectangle calculates the area of a rectangle
func Rectangle(base, height float64) float64 {
	return base * height
}

// private function
func _TriangleEq(base, height float64) float64 {
	return (base * height) / 2
}
