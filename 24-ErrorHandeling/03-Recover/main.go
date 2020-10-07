//explaining defer
package main

import "fmt"

func main() {
	b()
}

//here when we call func b(), the process is like this:
//i=0 > defer is called, so it will be execute at the end
//i=1 > defer is called, so it will be execute at the end
//i=2 > defer is called, so it will be execute at the end
//i=3 > defer is called, so it will be execute at the end
//we are at the end of the function, no the defer commands should run (we triggered 4 defer). so:
// the last refer will be execute first, so the output would be 3,2,1,0

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
