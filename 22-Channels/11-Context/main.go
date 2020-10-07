/*In Go servers, each incoming request is handled in its own goroutine. Request handlers often start
additional goroutines to access backends such as databases and RPC services.
The set of goroutines working on a request typically needs access to request-specific values
such as the identity of the end user, authorization tokens, and the request's deadline.
When a request is canceled or times out, all the goroutines working on that request should exit quickly
so the system can reclaim any resources they are using.
At Google, we developed a context package that makes it easy to pass request-scoped values,
cancelation signals, and deadlines across API boundaries to all the goroutines involved
in handling a request. The package is publicly available as context.
This article describes how to use the package and provides a complete working example.
Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

Incoming requests to a server should create a Context, and outgoing calls to servers should
accept a Context. The chain of function calls between them must propagate the Context,
optionally replacing it with a derived Context created using WithCancel, WithDeadline,
WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.

simply, when a process is killed, all Goroutins related to that process should be killed
*/

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context error:\t", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("---------------------")

	//WithCancel func(parent Context) (ctx Context, cancel CancelFunc)
	//WithCancel returns a copy of parent with a new Done channel.
	//The returned context's Done channel is closed when the returned cancel function is called or
	//when the parent context's Done channel is closed, whichever happens first.
	//Canceling this context releases resources associated with it, so code should call cancel
	//as soon as the operations running in this Context complete.

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("Context:\t", ctx)
	fmt.Println("Context error:\t", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("Cancle:\t", cancel)
	fmt.Printf("cancel type: %T\n", cancel)
	fmt.Println("---------------------")

	cancel()
	fmt.Println("Context:\t", ctx)
	fmt.Println("Context error:\t", ctx.Err())
	fmt.Printf("context type: %T\n", ctx)
	fmt.Println("Cancle:\t", cancel)
	fmt.Printf("cancel type: %T\n", cancel)
	fmt.Println("---------------------")

}
