//in go interfaces can not implement the other interfaces or extend them because inhiritence is not supported
//instead we can create a new interface by merging two or more interfaces, we call this interface embedding

//example4
package main

import (
	"fmt"
	"math"
)

//struct--------------------------------
type cube struct {
	edge  float64
	color string
}

//Methods-------------------------------
func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

func (c cube) getcolor() string {
	return c.color
}

//interfaces-----------------------------
type shapes interface {
	area() float64
}

type object interface {
	volume() float64
}

//declaration an interface means giving methods to it
//here we are merging all interfaces in one interface
//we can also call the other interfaces in it
//when we include an interface in another interface, we actually add all its method to the new interface
//so in this interface, geometry embed all the methods from shapes and volume and has its own method as well
type geometry interface {
	shapes
	object
	getcolor() string
}

//Functions--------------------------------------
//as you see here we don't have to call all methods in the interface. and here we only called area and volume
//so for any given interface in the function, we can return only the methods that we want
//like what you do with fmt package for instance
func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v

}

func main() {
	c := cube{edge: 2}
	a, v := measure(c)
	fmt.Printf("Area :%v, volume: %v\n", a, v)
}

//more explaintion:
//normally what we acheived in this example:
//we wanted to use one function to call all methods
//for that we need to have an interface to be able to call all methods
