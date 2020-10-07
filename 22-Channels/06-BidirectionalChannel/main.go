//we can convert a channel to two sub-channels (one receive only and one send only)
//we can use them in two different functions

package main

import "fmt"

func main() {

	//define a bidirectional
	c := make(chan int)

	go sendOnly(c)
	recieveOnly(c)
}

//semi-directional just send
func sendOnly(c chan<- int) {
	c <- 55
}

//semi-directional just receive
func recieveOnly(c <-chan int) {
	fmt.Println(<-c)
}
