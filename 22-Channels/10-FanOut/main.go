package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	p := make(chan int)
	job := make(chan int)
	go populate(p)
	go fanOut(p, job)
	for v := range job {
		fmt.Println(v)
	}
}

func populate(p chan int) {
	for i := 0; i < 100; i++ {
		p <- i
	}
	close(p)
}
func fanOut(p, j chan int) {
	var wg sync.WaitGroup
	wg.Add(1)
	//we define maximum gorouting that we want them to process simultenously
	const GoroutingMax = 10
	for i := 0; i < GoroutingMax; i++ {
		go func() {
			for v := range p {
				func(v2 int) {
					//this function is simulating a time consuming function
					j <- TimeConsumingTask(v)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(j)
}

//TimeConsumingTask ...
func TimeConsumingTask(v2 int) int {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return v2 + rand.Intn(1000)
}
