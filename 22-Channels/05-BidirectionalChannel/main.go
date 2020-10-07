package main

import "fmt"

func main() {

	c := make(chan<- int)
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	e := make(<-chan int)
	go func() {
		e <- 42
	}()

	fmt.Println(<-e)
}
