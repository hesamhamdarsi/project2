//asserting and type switch

//example3
package main

import (
	"fmt"
	"math"
)

type square struct {
	shape string
	high  float64
	width float64
}
type circle struct {
	shape   string
	reduies float64
}

//method for calculating area of square
func (s square) area() float64 {
	return s.high * s.width
}

//method for calculating primeter of square
func (s square) primeter() float64 {
	return (s.high + s.width) * 2
}

//method for calculating area of ciscle
func (r circle) area() float64 {
	return r.reduies * r.reduies * math.Pi
}

//method for calculating primeter of circle
func (r circle) primeter() float64 {
	return r.reduies * 2 * math.Pi
}

type shapes interface {
	area() float64
	primeter() float64
}

//this time I use type to print out the shape name
//s.(type) is asserting. it means we go one level down and access to the value of one level downer
func calculator(s shapes) {
	switch v := s.(type) {
	case circle:
		fmt.Println("Shape is : ", v)
	case square:
		fmt.Println("Shape is : ", v)
	}
	fmt.Println(s.area())
	fmt.Println(s.primeter())

}

func main() {
	square1 := square{
		shape: "circle",
		high:  3,
		width: 5,
	}
	circle1 := circle{
		shape:   "square",
		reduies: 2.5,
	}

	calculator(square1)
	calculator(circle1)
}
