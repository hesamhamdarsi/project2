//sending and returning values from different channels
//i want to return ood digits to a slice
//even digits to a slice
//quit result to an int
//select is only used in channels
package main

import "fmt"

func main() {

	evenList := []int{}
	oddList := []int{}
	Status := 0

	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	go send(odd, even, quit)

	oddList, evenList, Status = receive(odd, even, quit)
	fmt.Println(oddList)
	fmt.Println(evenList)
	fmt.Println(Status)

}

func send(o, e, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//we will use this q when we want to call our channels
	// this q would be the last number in the channel, so when we are reading, when we get q that means
	//we are in the last of the channel block. so we should return. otherwise, we get fatal error
	q <- 0
}

func receive(o, e, q <-chan int) ([]int, []int, int) {
	evenList := []int{}
	oddList := []int{}
	Status := 0
	for {
		select {
		//here we are take the value of the chanel o and put that in another variable
		case v := <-o:
			oddList = append(oddList, v)
		case v := <-e:
			evenList = append(evenList, v)
		//here we say when you reached to q, you should exit. I am exiting here with the slices result
		//if you had not any result value to return, you should simply use "return" without any variable
		//this would cause the loop to be stopped. otherwise, you will get the fatal error
		//because you didn't close the channel and select loop is still waiting to receive value
		case v := <-q:
			Status = v
			return oddList, evenList, Status
		}
	}

}
