package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error check 1:\t", ctx.Err())
	fmt.Println("Number of Goroutin at the begining:\t", runtime.NumGoroutine())
	fmt.Println("....")

	go func() {
		n := 0
		for {
			select {
			//Done() function of context use only for channels
			//it says, whenever you get the cancel signal, shut this process down
			//so the gorouting processes would be done.
			//in this example we've used WithCancel and so we've called cancel from that package.
			//we can use WithDeadline,WithTimeout, or WithValue as well to cancel the processes
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working ", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error check 2:\t", ctx.Err())
	fmt.Println("Number of Goroutin right before cancel:\t", runtime.NumGoroutine())

	fmt.Println("....")
	fmt.Println("about to cancel")
	//here we call Cancel()
	cancel()

	fmt.Println("....")
	fmt.Println("canceled context")
	time.Sleep(time.Second * 2)
	fmt.Println("Error check 3:\t", ctx.Err())
	fmt.Println("Number of Goroutin After cancel:\t", runtime.NumGoroutine())
}
