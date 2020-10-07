//error is a type
//we can create our own error type
//as it's a type of interface

package main

import (
	"errors"
	"fmt"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		fmt.Println(err)
	}
}
func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("not matched number")
	}
	return a, nil
}
