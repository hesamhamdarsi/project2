//in this example we take 2 string, and put their value in a channel seperately and then put the
//result of that channel in another channel

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := fanIn(boring("Hesam"), boring("Ali"))
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
}

func boring(msg string) <-chan string {
	var output = make(chan string)
	go func() {
		//here we have an unlimited loop, it is running by another Goroutin. buw how we have return?
		//actually return is from the channel which is kind of buffering data
		//the new data is still pushing there but at the same time we can return output
		//because it is actually an address in the memory
		//as Golang is a gc language (garbage cleaner) the memory wont be full and the allocated memory
		//is fixed
		//as we want this loop to work, we don't close the channel
		for i := 0; ; i++ {
			output <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return output
}

//here we say the output should be a read only channel
func fanIn(msg1, msg2 <-chan string) <-chan string {
	var c = make(chan string)
	go func() {
		for {
			//here we say the output of chnnel names msg1 into the channel c
			c <- <-msg1
		}
	}()
	go func() {
		for {
			c <- <-msg2
		}
	}()
	return c
}
