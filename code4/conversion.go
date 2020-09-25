package main

import (
	"fmt"
)

func main() {

	s := "My name is Hesam"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	//converting string to slice of bytes (asci codes)
	ConvertedToByte := []byte(s)
	fmt.Println(ConvertedToByte)
	fmt.Printf("%T\n", ConvertedToByte)
	//printing UTF-8 asci of each character inside the string
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}
}
