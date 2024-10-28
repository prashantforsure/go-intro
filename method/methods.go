package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

type Rectangle struct {
	length, width float64
}

// Method with a pointer receiver to modify width
func (r *Rectangle) Resize(newLength, newWidth float64) {
	r.length = newLength
	r.width = newWidth
}

// Method with a value receiver to calculate area
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func main() {
	c := Circle{
		radius: 5,
	}
	fmt.Printf("area: %.2f\n", c.Area())
    
	
	rect := Rectangle{length: 5, width: 10}
    fmt.Printf("Original Area: %.2f\n", rect.Area())

    // Resize modifies the original rectangle dimensions
    rect.Resize(8, 12)
    fmt.Printf("New Area after resize: %.2f\n", rect.Area())
}
