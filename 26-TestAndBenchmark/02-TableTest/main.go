//1- you need to create a file with name of "...._test.go" in the same directory as your package is
//2-run one/more func with a signature of "func TestXxx("t *testing.T") for any function you want to test
//		t *testing.T means a pointer to type T from package testing
//3- go .test

package main

import "fmt"

func main() {
	fmt.Println("sum of 2+3= ", sum(2, 3))
	fmt.Println("sum of 3+4= ", sum(3, 4))
	fmt.Println("sum of 7+8= ", sum(7, 8))
}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
