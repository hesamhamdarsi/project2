// we want to calculate circle and squre area
// first we make a struct type for each of them
// then we need to make some methods to get the exact struct type
//this method name is "area()" which is supposed to work for both square and circle
//when it get the instance of circle, then it could have access to the defined raduis
//when it get the instance of square,  then it could have access to the defined high and width
//as output we send out the calculated area
//now we make an interface to be a single gateway for both methods. so the user only call that interface for all shapes
//then we can make a function that get shape as an entry and return the result
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

func (s square) area() float64 {
	return s.high * s.width
}
func (r circle) area() float64 {
	return r.reduies * r.reduies * math.Pi
}

type shape interface {
	//this should totally match with method. when method is area() float64, the samething should be here
	area() float64
}

func calculator(c shape) {
	fmt.Println(c.area())
}

func main() {
	square1 := square{
		high:  3,
		width: 5,
	}
	circle1 := circle{
		reduies: 2.5,
	}

	calculator(square1)
	calculator(circle1)
}
