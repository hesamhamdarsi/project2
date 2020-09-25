package main

import (
	"fmt"
	"runtime"
)

var x bool

func main() {

	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 42
	b := 30
	//print if
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a <= b)

	//print out the system architecture and OS
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
