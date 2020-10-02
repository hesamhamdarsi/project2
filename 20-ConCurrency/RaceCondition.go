// https://blog.regehr.org/archives/490
// here we want to check
// for one test use go run
// for second test use go run -race > this one shows if there is ant data race. we need to fix data race later
//the counter should return 100 at the end of the execute, but as we have race condition, it wount be 100
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
	v := Counter
	//Gosched() tell the thread if you want to do another process you can do it
	//it's actually yeild. yeild in computer process management means:
	//when a shared reource (like a variable like counter) want to be accessible from different threads
	// each thread first read this variable, then yeild, means announce the other threads that
	//if you want you can use this variable, I'm done with that
	//then it will increase it's internal counter
	runtime.Gosched()
	v++
	Counter = v
	wg.Done()
}
