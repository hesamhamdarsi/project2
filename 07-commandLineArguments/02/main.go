//example:
//getting a value from the cli, check if it's only one value and if it is int
//then do something with that

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if args := os.Args; len(args) > 2 {
		fmt.Println("enter only one digit...")
	} else if len(args) != 2 {
		fmt.Println("you need to enter a digit...")
	} else if km, err := strconv.ParseInt(args[1], 10, 32); err != nil {
		fmt.Println("only digits are accepted...")
	} else {
		fmt.Println("the speed in 100 km = ", 100*km)
	}
}

//two points here:
//you can combine initial step and if statement together in a line.
//for instance "if args := os.Args; len(args) > 2 " will first make a variable called args
//and asign the value to it, then you can use this value for condition checking
//point2:
//os.Args has a value in args[0] as well. but that value is the entire command itself
//so you shouldn't call that
