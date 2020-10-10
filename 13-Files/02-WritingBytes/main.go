//we have two ways for that
//using os package
//using ioutil package
//ioutil is simpler and have more features but its makes your file heavy. so do not use that if you don't have to

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	//method 1
	file, err := os.OpenFile("myFile.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//you need to give it he byte of slice
	byteSlice := []byte("I have an Idea to work on")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Byte Written: %d\n", byteWritten)

	//##############################################################
	fmt.Println(strings.Repeat("#", 20))
	//##############################################################

	//method2
	//it will truncate, create, open and close the file automatically
	//it will create bytewritten automatically
	//you need to give it the byte of slice
	byteSlice = []byte("I have my freinds on my back")
	err = ioutil.WriteFile("myFile2.txt", byteSlice, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
