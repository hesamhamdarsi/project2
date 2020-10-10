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
