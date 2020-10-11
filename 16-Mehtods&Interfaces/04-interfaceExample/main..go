/*interface is collection of method signatures that an object (which is a named type) can implement
interfaces define the behavior of an object and can acheive polymorphism(the condition of occurring in several different forms.)
when two methods have the same logic, for instance both of them calculate the area and/or primeter
we can use an interface for them. in this way, instead of creating several seperate functions
for instance for circle, rectangle, etc (all shapes) we will create only one function
the way we handel this:
1- we create different methods which are for different shapes
	this methds are calculating area and primeter of those shapes
2- here we could have two approach:
	2-1- using different functions (example1) for calculating each shapes area and primeter
		 using this functions, we path the values to the function, and function will call the method
	2-2- using interface (example2):
		 the interface contains only the signatre of the methods
		 the signature of a method is its name, input parameters and return values
		 interface is also a type, so we need to create this type
*/

//example2
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
// in the primeter, i also return the name of the shape
func (s square) primeter() (float64, string) {
	return (s.high + s.width) * 2, s.shape
}

//method for calculating area of ciscle
func (r circle) area() float64 {
	return r.reduies * r.reduies * math.Pi
}

//method for calculating primeter of circle
func (r circle) primeter() (float64, string) {
	return r.reduies * 2 * math.Pi, r.shape
}

//creating an interface for common logic
//the type of interface is dynamic(polymorphism), that means when we call it as a circle it became a
//circle and when we call it as a rectangle, it became a rectangle
//so in the next example we use the type switching(example3)
type shapes interface {
	//method name (parameters if there is any) returned values if there is any
	area() float64
	primeter() (float64, string)
}

//this time, our function do not get a circle or square. but it takes any shape
func calculator(s shapes) {
	pri, name := s.primeter()
	fmt.Println(name)
	fmt.Println(s.area())
	fmt.Println(pri)
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
