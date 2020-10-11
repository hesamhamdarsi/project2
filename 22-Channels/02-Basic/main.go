//we have another type of channel which is buffer channel that do not need a allocated seperate process
//this is like a buffer and you can put there values as much as you want and get them back in a row

//example2:
//this works as we run a seperate Goroutin for that

package main

import (
	"fmt"
)

func main() {

	//creating a channel (which is a type) of int
	c := make(chan int)
	//the best practice is we close the channel after all tasks are done.
	//as here the main function is the function that will finish after all other functions, we close the channel after it ends
	defer close(c)

	//putting value to this channel should be through a seperate Goroutin
	go func() {
		c <- 42
	}()

	//i am putting another value to channel buffer using another seperate Goroutin
	go function1(23, c)

	//calling the channel (buffered values)
	fmt.Println(<-c)
	fmt.Println(<-c)

}

func function1(n int, ch chan int) {
	ch <- n
}
