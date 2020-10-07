//fan out means passing processes to different channels as much as possible
//fan in means collecting these reults to only one channel
//---------------------------
//in this example we take 2 channels and put their value in the second channel

package main

import (
	"fmt"
	"sync"
)

func main() {

	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanIn)

	//the loop is going to be blocked here until "fanIn" channel is closed
	for v := range fanIn {
		fmt.Println(v)
	}

}

func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		//the loop is going to be blocked here until "even" channel is closed
		for v := range even {
			fanIn <- v
		}
		wg.Done()
	}()
	go func() {
		//the loop is going to be blocked here until "odd" channel is closed
		for v := range odd {
			fanIn <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanIn)
}
