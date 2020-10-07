//bidirectional means we define the channel to be able to put and get vlue in that
// we can define channel in a way that only put in that, or only get from ot(semi directional)
// for example, we make two progrmas, in one of them the clients will pass the value and on the other
// the clients will get the value back only

package main

import (
	"fmt"
)

func main() {

	//only send to channel
	//you can not get the value back in this function
	c := make(chan<- int, 2)
	fmt.Printf("%T\n", c)
	c <- 42
	//you can not receive the value from it
	//fmt.Println(<-c)

	e := make(<-chan int, 1)
	fmt.Printf("%T\n", e)
	//you can not put the value in it
	//e <- 40
	//fmt.Println(<-e)
}
