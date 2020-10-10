package main

import "fmt"

//expression means declaring function to a variable
//outside
var f = func(x int) {
	fmt.Printf("my number is: %+v\n", x)
}

//expression means declaring function to a variable
//inside
func main() {

	g := func(x int) {
		fmt.Printf("my number is: %+v\n", x)
	}
	f(22)
	g(33)
}
