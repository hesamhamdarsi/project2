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
	fmt.Println("")
	for i, v := range s {
		fmt.Println(i, v)
	}
	for i, v := range s {
		//print hex
		fmt.Printf("the %d'th value is %#x\n", i, v)
		//bianry
		fmt.Printf("the %d'th value is %b\n", i, v)
		//base-10
		fmt.Printf("the %d'th value is %d\n", i, v)
		fmt.Println("")
	}
}
