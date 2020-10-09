package main

import (
	"fmt"
	"strings"
)

func main() {

	//to give the initial size and capacity to a slice
	// make([]type, size, capacity)
	myComposite := make([]int, 10, 100)

	fmt.Println(myComposite[:])
	fmt.Println(len(myComposite))
	fmt.Println(cap(myComposite))
	myComposite[0] = 1
	myComposite[4] = 5
	myComposite[6] = 8
	fmt.Println(myComposite[:])
	fmt.Println(len(myComposite))
	fmt.Println(cap(myComposite))
	myComposite = append(myComposite, 7, 8, 9)
	fmt.Println(myComposite[:])
	fmt.Println(len(myComposite))
	fmt.Println(cap(myComposite))

	fmt.Println(strings.Repeat("#", 20))
	//copy a slice to another slice:
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Println(slice2)

}
