package main

import "fmt"

func main() {

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

}
