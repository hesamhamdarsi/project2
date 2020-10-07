//bidirectional means we define the channel to be able to put and get vlue in that
// we can define channel in a way that only put in that, or only get from ot(semi directional)
// for example, we make two progrmas, in one of them the clients will pass the value and on the other
// the clients will get the value back only

package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 2)
	fmt.Printf("%T\n", c)
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)

}
