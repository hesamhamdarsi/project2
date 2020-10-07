package main

import (
	"fmt"
)

//Using fmt.Errorf
//fmt.Errorf will run a errors.New, but the difference is you can pass a value in that as well
func main() {
	_, err := sqrt(-10)
	if err != nil {
		fmt.Println(err)
	}
}
func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("not matched number for '%+v'", a)
	}
	return a, nil
}
