//create and write to a file
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//this function will create a file and return the fileAddress and error in there is any
	Myfile, err := os.Create("Myfile.txt")
	if err != nil {
		fmt.Println(err)
		//this return says exit from this function and close the program
		return
	}
	//we will close the file to save the resource after finishing the main func
	defer Myfile.Close()
	// in order to write into the file, we need a reader interface type and a writer interface
	//the reader interface here is provided by newReader function
	// the writer interface is going to be the file
	reader := strings.NewReader("My name is Hesam")
	io.Copy(Myfile, reader)

}
