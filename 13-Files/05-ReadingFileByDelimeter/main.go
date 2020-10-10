//reading file line by line or by any other delimeter

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//first we need to open the file
	file, err := os.Open("myFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//then we define a reader interface from bufio package
	scanner := bufio.NewScanner(file)

	//we check the output of the scan() function of this interface
	//if it is true, it means we have line, else means there is no line or no more line
	//the default delimeter is line by line. if we want to use word by word OR even char be char, we
	//should define the splitter
	//scanner.Split(bufio.ScanWords)
	//scanner.Split(bufio.ScanRunes)
	output := scanner.Scan()
	err = scanner.Err()

	//if its flase, that means either we have erro or there is no line
	if output == false {
		if err == nil {
			log.Println("Scan was completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("The first line is: ", scanner.Text())
	//fmt.Println("The first line is: ", scanner.Bytes())  >> this will retuen slice of bytes

	//now we want to print all lines, the above example's output was only for one line
	//anytime we call scanner.Scan, the pointer will go to the next line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err = scanner.Err(); err == nil {
		log.Println("Scan was completed and reached EOF")
	} else {
		log.Fatal(err)
	}
}
