package main

import (
	"fmt"
)

//fallback means passing a func as an argument
//example: return odd numbers

func main() {
	//we have a slice of int numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// we send this slice to the even function
	//even function is responsible to pick up the odd digits and return another slice
	//as we want to send slice, we need to use ... to indicate that
	result := even(numbers...)

	//the result is also slice, so we need to use ... to pass it to sum function
	x := sum(result...)
	fmt.Println(x)
}

//the argument is slice, so we have to define it like: func <identifier> (instance ...<type>) retrrned_type
func sum(ii ...int) int {
	var summary int
	for _, v := range ii {
		summary += v
	}
	return summary
}

//we are returning a slice, so we need to indicate it via "[]int" as it is slice of int
func even(ii ...int) []int {
	var oddNumbers []int
	for _, v := range ii {
		if v%2 != 0 {
			oddNumbers = append(oddNumbers, v)
		}
	}
	return oddNumbers
}
