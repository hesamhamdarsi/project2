package main

import (
	"errors"
	"fmt"
)

//ErrorOfDigit ...
//here we will make our own error as a variable
var ErrorOfDigit = errors.New("not matched number")

func main() {
	_, err := sqrt(-10)
	if err != nil {
		fmt.Println(err)
	}
}
func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrorOfDigit
	}
	return a, nil
}
