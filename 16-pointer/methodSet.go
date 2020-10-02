//method set is simply the methods that are attached to a type
//methods could have different types as reciever:
//1- non-pointers
//2- pointers
//a method which is set to accept non-pointer types, can accept both pointer and non-pointer types
//a method which is set to accept only pointer types, can only accept the pointer type unles we call it directly

package main

import (
	"fmt"
)

type computers struct {
	architecture string
	model        string
	price        float64
}

func (c computers) color() string {
	if c.price > 120 {
		return "black"
	} else {
		return "white"
	}
}

type printcolor interface {
	color() string
}

func barcode(p printcolor) {
	fmt.Printf("the laptop color is %+v\n", p.color())
}

func main() {

	myComputer := computers{
		architecture: "amd64",
		model:        "ausus2020",
		price:        125.5,
	}

	fmt.Printf("myComputer=%+v\n", myComputer)
	fmt.Println("&myComputer= ", &myComputer)
	fmt.Println("&myComputer.model= ", &myComputer.model)
	//both myComputer and &mycomputer could be shiped to the method
	//but if methods were defined this way: "func (c *computers) color() string", then it only accepts pointer
	// *computers means a pointer of computer type
	barcode(myComputer)
	barcode(&myComputer)

}
