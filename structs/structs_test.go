package structs

import "testing"

func TestPerimeter(t *testing.T) {
	expected := 40.0
	rectangle := Rectangle{10.0, 10.0}
	result := rectangle.Perimeter()

	if result != expected {
		t.Errorf("Got: %.2f, but expected: %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, expected: 100},
		{name: "Circle", shape: Circle{10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Area()
			if result != tt.expected {
				t.Errorf("%#v got: %g, but expected: %g", tt.shape, result, tt.expected)
			}
		})

	}
}
