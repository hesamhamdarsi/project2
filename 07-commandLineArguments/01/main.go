//os.Args get a slice of strings from the command line.package main
//notice: its always string, so if you want to get an interger form the CLI, you need to
//get it as string and convert it to int

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	//to get the first argument
	var result, _ = strconv.ParseInt(os.Args[1], 10, 64)
	fmt.Println(result + 2)

	// to get all of them
	for _, v := range os.Args {
		fmt.Println(v)
	}
	/*
		go run main.go 10 12 13
		12
		/tmp/go-build091323392/b001/exe/main
		10
		12
		13
	*/

	//OR

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
