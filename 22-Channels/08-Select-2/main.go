//sending and returning values from different channels
//i want to return ood digits to a slice
//even digits to a slice
//this time we want to use bool for quite
package main

import "fmt"

func main() {

	evenList := []int{}
	oddList := []int{}
	Status := true

	odd := make(chan int)
	even := make(chan int)
	quit := make(chan bool)

	go send(odd, even, quit)

	oddList, evenList, Status = receive(odd, even, quit)
	fmt.Println(oddList)
	fmt.Println(evenList)
	fmt.Println(Status)

}

func send(o, e chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//we are going to close all channels
	close(e)
	close(o)
	close(q)
}

func receive(o, e <-chan int, q <-chan bool) ([]int, []int, bool) {
	evenList := []int{}
	oddList := []int{}
	for {
		select {
		//if (<-o) had returned two outputs, we used for instance "case i , v"
		//when you close the channel, it will return 2 values, the first one is the value in the block
		//and the second one is status
		//when we still have value in our block, the status is going to be true.
		//when values are finished, the status would be false
		case v, status := <-o:
			if status {
				oddList = append(oddList, v)
			}
		case v, status := <-e:
			if status {
				evenList = append(evenList, v)
			}
		//here we say when you reached to q, you should exit. I am exiting here with the slices result
		//if you had not any result value to return, you should simply use "return" without any variable
		//this would cause the loop to be stopped. otherwise, you will get the fatal error
		//I dont need to check the status of quit because it's definatly flase, so I used _
		case v, _ := <-q:
			lastStatus := v
			return oddList, evenList, lastStatus
		}
	}

}
