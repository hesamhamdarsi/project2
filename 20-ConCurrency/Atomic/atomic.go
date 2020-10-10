// we've used go run -race > this shown there is data race. we need to fix data race later
// we do it via mutex function. though this function, thread will lock a variable until it finish its job
// then this variable would be accessible for the other threads
// the counter will return 100 at the end of the execute.
// this is what we call Synchronization
// we have another method to do synchronization, and that's using "atomic" package
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Counter ...
var Counter int64 = 0

const gr = 100

var wg sync.WaitGroup

func main() {

	wg.Add(gr)

	fmt.Println("CPU Cores: \t\t", runtime.NumCPU())
	fmt.Println("GoRoutin: \t\t", runtime.NumGoroutine())
	for i := 0; i < gr; i++ {
		go count()
		fmt.Println("GoRoutin: \t\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter: \t\t", Counter)
	fmt.Println("GoRoutin: \t\t", runtime.NumGoroutine())
}

func count() {
	//this will give the value from address and increase it one unit
	//but this only accept Delta (increment and decrement).
	atomic.AddInt64(&Counter, 1)
	//we can remove runtime.Gosched() from this code as we are locking the value through Mutex
	//we've used that in previous examples to make some delys for the other threads to be run
	runtime.Gosched()
	//fmt.Println("Counter: \t ", atomic.LoadInt64(&Counter))
	wg.Done()
}
