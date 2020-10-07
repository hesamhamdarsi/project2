//if you see documentation about the type error you see it is identified by:
/*
type error interface{
	Error() string
}
*/
//that means if we have a type like struct that has a method with the same signature like above
//then any value of that type also implement the error interface
package main

import (
	"fmt"
)

//MyErrorStruct ...
type MyErrorStruct struct {
	item1 string
	item2 int
	err   error
}

//the method
func (e MyErrorStruct) Error() string {
	return fmt.Sprintf("you have an error with %v and level of %d related to: %v", e.item1, e.item2, e.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		fmt.Println(err)
	}
}
func sqrt(a float64) (float64, error) {
	if a < 0 {
		errToPrint := fmt.Errorf("not matched number for '%+v'", a)
		return 0, MyErrorStruct{"type of Hesam's Error", 12, errToPrint}
	}
	return a, nil
}
