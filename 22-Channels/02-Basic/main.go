//putting on the channel and taking of the cahnnel should be in parallel. that means exchange data
//should be simultenously. that means the process of putting on the channel and taking off the channel
//should be running together. and that means both of them need to be run in seperate threads (Goroutins)
//actually channels allows pass value between Goroutins
//channels are the better way than waitgroups. because in wait groups several goroutins access the same
//memory address which value is stored in. but via channels they are in different memory location
//and via channels Goroutins can share these values between eachothers

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

	//putting value to this channel should be through a seperate Goroutin
	go func() {
		c <- 42
	}()

	//calling the channel
	fmt.Println(<-c)
}
