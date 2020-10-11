//channels are used to communication between goroutins
//data which is sending to the channel and receiving should be from the same type always
//channels are addresses. so they are like pointers
//channels are bi-directional. you can send and receive in them.
// SEND:     c <- 10
//Receive:   num := <- c
// to close a channel we use Close(c). so sends will panic and recieves will get zero after that
// if you work with a channel normally you will get the error "deadlock!". its because channels work only with goroutins

//putting on the channel and taking of the cahnnel should be in parallel. that means exchange data
//should be simultenously. that means the process of putting on the channel and taking off the channel
//should be running together. and that means both of them need to be run in seperate threads (Goroutins)
//actually channels allows pass value between Goroutins
//channels are the better way than waitgroups. because in wait groups several goroutins access the same
//memory address which value is stored in. but via channels they are in different memory location
//and via channels Goroutins can share these values between eachothers
//if we don't do these process simultenously, the channels would be blocked

//we have another type of channel which is buffer channel that do not need a allocated seperate process
//this is like a buffer and you can put there values as much as you want and get them back in a row
//through buffer channels, there is no more blocking
//but if we  define the number of buffer correctly, the program will work. otherwise, wont. so we
//dont recommend to use that

//example1:
//this wont work because channel need to be run in a seperate process and in paralel, otherwise
//the value can not be passed between the sender and reciever

package main

import (
	"fmt"
)

func main() {

	//initializing channel
	c := make(chan int)
	c <- 42

	fmt.Println(<-c)
}
