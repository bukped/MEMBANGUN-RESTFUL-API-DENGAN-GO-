package main

import "fmt"

// Fungsi
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Metode
type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}
func main() {
	greet("John")
	rectangle := Rectangle{length: 5, width: 3}
	fmt.Println("Area:", rectangle.area())
}
