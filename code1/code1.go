package main

import (
	"fmt"
)

func main() {
	fmt.Println("before loop and fucntion")
	my_func()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("After loop and function")
}

func my_func() {
	fmt.Println("after fucntion and before loop")
}
