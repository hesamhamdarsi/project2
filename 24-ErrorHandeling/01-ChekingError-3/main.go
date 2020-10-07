//Read from an exsisting File
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//this function will read from an existing file and return the fileAddress and error if
	//there is any. for instance if file is not available
	Myfile, err := os.Open("Myfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer Myfile.Close()

	// to read a file we use ReadAll function from io/ioutil package. it return Slice of bytes
	//so we need convert those bytes to string
	bs, err := ioutil.ReadAll(Myfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))

}
