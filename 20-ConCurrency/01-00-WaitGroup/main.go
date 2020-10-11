//goroutin VS thread : https://www.geeksforgeeks.org/golang-goroutine-vs-thread/
//OS threads are scheduled by the kernel but goroutings are schedule by own go schedulers using a technic called m:n scheduling
//it multiplex (schedule) m go routing on n os threads

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var x int = 0
var y int = 0

//if we dont define wg here and if we want to define it in the main function, we should pass it's pointer to the go funs()
var wg sync.WaitGroup

func main() {

	fmt.Printf("OS:\t\t%+v\n", runtime.GOOS)
	fmt.Printf("arch:\t\t%+v\n", runtime.GOARCH)
	fmt.Printf("CPU:\t\t%+v\n", runtime.NumCPU())
	//the number of active goroutins (for now its only one for the main function)
	fmt.Printf("Goroutings:\t%+v\n", runtime.NumGoroutine())
	//the number of OS active threads. the default value is number of CPU cores
	//but you can change it to any number you want (it is the m part in m:n scheduling)
	//0 means return the current setting. if we want to change it, we'll use another digit in it
	fmt.Printf("GoMaxproc:\t%+v\n", runtime.GOMAXPROCS(0))
	fmt.Println("X=\t\t", x)

	//here we say two things we will be waited for until we get both of them done() function
	//if we don't use Waitgroup, we have to make kind of delay in the main gorouting so that the second goroutin finish its job
	//otherwise the second one do not have enough time to finish its job
	wg.Add(2)
	go loop1()
	go loop3()
	loop2()

	fmt.Printf("Goroutings:\t%+v\n", runtime.NumGoroutine())
	//we wait here until we get done() signals
	wg.Wait()
	fmt.Printf("Goroutings at the end of main:\t%+v\n", runtime.NumGoroutine())
}

func loop1() {
	fmt.Println("we are in function 1 and wait for 5 Second")
	time.Sleep(time.Second * 10)
	fmt.Println("functione 1 is done")
	fmt.Printf("Goroutings at the end of function 1:\t%+v\n", runtime.NumGoroutine())
	//the first done()
	wg.Done()

}
func loop2() {
	fmt.Println("we are in function 2 and wait for 8 Second")
	time.Sleep(time.Second * 13)
	fmt.Println("functione 2 is done")
	fmt.Printf("Goroutings at the end of function2:\t%+v\n", runtime.NumGoroutine())
}
func loop3() {
	fmt.Println("we are in function 3 and wait for 11 Second")
	time.Sleep(time.Second * 13)
	fmt.Println("functione 3 is done")
	fmt.Printf("Goroutings at the end of function 3:\t%+v\n", runtime.NumGoroutine())
	//the second done()
	wg.Done()
}
