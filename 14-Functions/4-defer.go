package main

import "fmt"

func main() {

	//defer: if we use defer, that means delay to call this function
	//the delay is until we finish the main() function here (or any function that we are in)
	//right after the open function is red, the defer functions will be execute
	//this is used to reduce the memory, because when we read the funtion, we'll close it and do not
	//take action on it until we finish the func that we are in and then we'll execute that
	defer delayInCall()
	defer delayInCall2()
	normallCall()
}
func delayInCall() {
	fmt.Println("defer")
}
func normallCall() {
	fmt.Println("normal")
}
func delayInCall2() {
	fmt.Println("defer2")
}
