//buffered writer will write in the file one time
//that means you maybe want to do a lot of process on data before writing that
//so you will buffer them and only when your work is done you will open the file and write into it
//it accepts both string and slice of bytes

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//first we open the file
	file, err := os.OpenFile("myFile.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//we'll make the asci codes (slices) or strins that we need (we can do that later)
	bs := []byte{97, 98, 99}
	st := "My name is Hesam Hamdarsi"

	//we make a nrew writer interface
	bufferWriter := bufio.NewWriter(file)

	//we write our ascis to the beffer in this step, then in the other steps we will add string as well
	byteWritten, err := bufferWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}

	//here we check how much bytes are written to the buffer by this slice
	fmt.Printf("Bytes Written to buffer(still no file)-slice: %v\n", byteWritten)

	//check how much capacity is remained in to the buffer. by default the buffer size is 4096 bytes
	availableBytesToWrite := bufferWriter.Available()
	fmt.Printf("Bytes remains in buffer: %v\n", availableBytesToWrite)

	//as we want a new line in the file before adding string we add the new line here
	st = "\n" + st

	//adding the string to the buffer
	byteWritten, err = bufferWriter.WriteString(st)
	if err != nil {
		log.Fatal(err)
	}
	//check how much bytes are written by this string
	fmt.Printf("Bytes Written to buffer(still no file)-string: %v\n", byteWritten)

	availableBytesToWrite = bufferWriter.Available()
	fmt.Printf("Bytes remains in buffer: %v\n", availableBytesToWrite)

	//check how much bytes are written in total
	storedDataOnBuffer := bufferWriter.Buffered()
	fmt.Printf("Bytes buffered total: %v\n", storedDataOnBuffer)

	//if we want to reset the buffer to make it empty before writing to the file, we use this command
	//bufferWriter.Reset(bufferWriter)

	//wite the buffer to the file and flush the buffer
	bufferWriter.Flush()
}
