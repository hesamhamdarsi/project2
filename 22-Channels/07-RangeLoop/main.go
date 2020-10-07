//we want to insert a bunch of values in the channel and call them back

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
	for i := 0; i < 100; i++ {
		c <- i
	}
	//we use close to close the channel after we write in it. this doesn't mean close the func.
	//this means we say we don't want to put more value in it
	close(c)
}

//semi-directional just receive
func recieveOnly(c <-chan int) {
	for v := range c {
		//for i := 0; i < 100; i++ {
		//if we had not use close(), we would get a fatal error here during the recieving, because
		// it had to wait for other entries in the channel. unless we use exactly for i := 0; i < 100; i++
		//this will bring the first 100 value that we have already in channel.
		//in this case, close won't be needed
		fmt.Println(v)
	}
}
