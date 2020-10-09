//breake and continue use only in for and switch statements
//but go to can be used anywhewre

package main

import (
	"fmt"
)

func main() {

	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	fmt.Println("the end")
}

//notice:
//you can not use the following goto:
/*
	go to next
	x := 5
next:
	someCommands....

because you can not jump to a label via go to if there is at least a new variable between
the goto and label. there should not be any new declarate variable
*/
