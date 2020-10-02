package main

import (
	"fmt"
	"runtime"
	"sync"
)

var x int = 0
var y int = 0

var wg sync.WaitGroup

func main() {

	fmt.Printf("OS:\t\t%+v\n", runtime.GOOS)
	fmt.Printf("arch:\t\t%+v\n", runtime.GOARCH)
	fmt.Printf("CPU:\t\t%+v\n", runtime.NumCPU())
	fmt.Printf("Goroutings:\t%+v\n", runtime.NumGoroutine())
	fmt.Println("X=\t\t", x)

	//here we say two things we will be waited for until we get both of them done() function
	wg.Add(2)
	go loop1()
	go loop3()
	loop2()

	fmt.Printf("Goroutings:\t%+v\n", runtime.NumGoroutine())
	//we wait here until we get done() signals
	wg.Wait()
	fmt.Println("X=\t\t", x)
	fmt.Println("Y=\t\t", y)
}

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		x = i
	}
	//the first done()
	wg.Done()

}
func loop2() {
	for j := 0; j < 10; j++ {
		fmt.Println("m")
	}
}
func loop3() {
	for j := 0; j < 10; j++ {
		fmt.Println(j)
		y = j
	}
	//the second done()
	wg.Done()
}
