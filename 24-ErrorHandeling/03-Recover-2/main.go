/*Panic is a built-in function that stops the ordinary flow of control and begins panicking.
When the function F calls panic, execution of F stops, any deferred functions in F are executed
normally, and then F returns to its caller. To the caller, F then behaves like a call to panic.
The process continues up the stack until all functions in the current goroutine have returned,
at which point the program crashes. Panics can be initiated by invoking panic directly.
They can also be caused by runtime errors, such as out-of-bounds array accesses.

Recover is a built-in function that regains control of a panicking goroutine.
Recover is only useful inside deferred functions. During normal execution, a call to recover will
return nil and have no other effect. If the current goroutine is panicking, a call to recover will
capture the value given to panic and resume normal execution.

Here's an example program that demonstrates the mechanics of panic and defer:

all of these simply means when we use recover(with defer) all process will be reversed until we reach
to the main function and pass the point that cause error and normally exit the program*/

//when you execute this program, it will tell you when recover just occured. it says 4.
//in this program that means when i was 4, we had panic

package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
