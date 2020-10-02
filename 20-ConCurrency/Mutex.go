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
)

//Counter ...
var Counter int = 0

const gr = 100

var wg sync.WaitGroup

//define mu variable
var mu sync.Mutex

func main() {

	wg.Add(gr)

	fmt.Println("CPU Cores: \t\t", runtime.NumCPU())
	fmt.Println("GoRoutin: \t\t", runtime.NumGoroutine())
	for i := 0; i < gr; i++ {
		go count()
		fmt.Println("GoRoutin: \t\t", runtime.NumGoroutine())
	}
	//fmt.Println("GoRoutin: \t\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Counter: \t\t", Counter)
	fmt.Println("GoRoutin: \t\t", runtime.NumGoroutine())
}

func count() {
	//use mu variable to luck the variable "Counter" right before gerping that
	//since now, no one can work on this variable
	//we can use any value we want in V to pass it to the counter. for instance: Counter = (v*10)+1
	mu.Lock()
	v := Counter
	//we can remove runtime.Gosched() from this code as we are locking the value through Mutex
	//we've used that in previous examples to make some delys for the other threads to be run
	runtime.Gosched()
	v++
	Counter = v
	//now we can unluck it
	mu.Unlock()
	wg.Done()
}
