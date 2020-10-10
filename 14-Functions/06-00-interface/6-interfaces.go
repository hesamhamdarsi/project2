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
	2-2- using interface (example2)
*/

//example1
package main

import (
	"fmt"
	"math"
)

type square struct {
	high  float64
	width float64
}
type circle struct {
	reduies float64
}

//method for calculating area of square
//here the behavior is area() >> calculating area (actually its the name that shows the behavior as well)
func (s square) area() float64 {
	return s.high * s.width
}

//method for calculating primeter of square
//here the behavior is primeter() >> calculating primeter (actually its the name that shows the behavior as well)
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

func circleCalculator(c circle) {
	//here we get the value from the main function, and attach the method to it (c.area())
	fmt.Printf("the area of circle is: %v\n", c.area())
	fmt.Printf("the primeter of circle is: %v\n", c.primeter())
}

func squareCalculator(s square) {
	fmt.Printf("the area of square is: %v\n", s.area())
	fmt.Printf("the primeter of square is: %v\n", s.primeter())
}

func main() {
	square1 := square{
		high:  3,
		width: 5,
	}
	circle1 := circle{
		reduies: 2.5,
	}

	circleCalculator(circle1)
	squareCalculator(square1)
}
