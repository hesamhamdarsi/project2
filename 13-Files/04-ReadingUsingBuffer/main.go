package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//creating a pointer to the file, using for the first and second method
	//no matter what method you are using, all of them read the file as slices of bytes
	file, err := os.Open("myFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//method1:
	//creating a buffer with size 2 bytes
	byteSlice := make([]byte, 2)

	numberOfBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("numver of bytes read: %v\n", numberOfBytesRead)
	//you need to use %s to convert byte to string
	fmt.Printf("data:  %s\n", byteSlice)

	//##############################################################
	fmt.Println(strings.Repeat("#", 20))
	//##############################################################

	//method2:
	//reading the whole file and accessing directly using ReadAll(you should spesify the pointer to file)
	//this need the file to be opened first
	file, err = os.Open("myFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("number of bytes read: %v\n", len(data))
	fmt.Printf("data:  %s\n", data)

	//##############################################################
	fmt.Println(strings.Repeat("#", 20))
	//##############################################################

	//method3:
	//reading the whole file and accessing directly using ReadFile(you should spesify file name as string)
	//this funciton handel opening and closing the file
	data2, err := ioutil.ReadFile("myFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("number of bytes read: %v\n", len(data))
	fmt.Printf("data:  %s\n", data2)

}
