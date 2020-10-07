//Error handling:
//Go do not accept the exception because it could have negative effect
//Erro is a type in Golang
//It is a type of interface, so any type that has method Error() attach to it is also an
//error package main

package main

import (
	"fmt"
)

func main() {

	var answer1, answer2, answer3 int
	fmt.Print("Your name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}
	fmt.Print("Your country: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}
	fmt.Print("Your age: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println("your answer is: ", answer1, answer2, answer3)
}
