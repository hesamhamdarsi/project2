//by default bufio use new line to finish the scan (if use press enter)

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//this first exmaple take only one line
	fmt.Print("please enter a text: ")
	scanner := bufio.NewScanner(os.Stdin)

	//here scanner is waiting for enter(new line) by default
	scanner.Scan()

	text := scanner.Text()
	bytes := scanner.Bytes()

	fmt.Println("text is: ", text)
	fmt.Println("bytes are: ", bytes)

	//how get multiple line or a file from the CLI?
	fmt.Print("please enter several line text and then issue quit: ")
	scanner = bufio.NewScanner(os.Stdin)

	//to put all texts in a slice
	allTexts := []string{}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "quit" {
			fmt.Println("Exit the function...")
			break
		}
		allTexts = append(allTexts, text)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for _, v := range allTexts {
		fmt.Println(v)
	}
}
