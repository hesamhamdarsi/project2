package main

import "fmt"

var numbers [10]int

//you can use init functions in orfer to initialize some variables, etc before running the main packge
func init() {
	for i := 0; i < 10; i++ {
		numbers[i] = i
	}
}

func main() {
	fmt.Printf("%d", numbers)

}
