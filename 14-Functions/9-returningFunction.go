package main

import (
	"fmt"
)

func main() {

	//returning a function means return a function as the result of another function
	//we can use anonymes function

	x := func1()
	fmt.Printf("%T\n", x)

	//after we assign output to a variable, this variable now is a function. so we can call it
	output := x()
	fmt.Println(output)

	//OR very shorter
	//x = func1() ---> x() = func1()+()
	fmt.Println(func1()())
}

//func <function_identifier> <returned type>
//func       func1()            func()int
func func1() func() int {
	return func() int {
		return 253
	}
}
