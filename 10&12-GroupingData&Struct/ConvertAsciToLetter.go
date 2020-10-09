//in golang, we don't have char, so if you want to get a char value and print it, you will
//see a number instead. that number is int32 or uint64
//we call them ascii codes
//notice: string packages will help us a lot in case of chars and strings

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	i := 0
	for {
		if i > 121 {
			break
		}
		if i < 64 {
			i++
			continue
		}
		i++
		//rune will get the unicode number
		fmt.Printf("%+v - ", rune(i))
		fmt.Printf("%c - ", i)
		fmt.Printf("%+c - ", rune(i))
		fmt.Printf("%c", int32(i))
		fmt.Println("")
	}
	fmt.Println("")

	//print an string characters
	//we have two methods:
	//1- manually
	//as in golang, we don't have char, so if you want to get a char value and print it we
	//see a number instead. that number is int32 or uint64
	//so we need to decode that number
	//the length wont be increased 1 by 1, becuase it is not char but it is bytes
	//so we get the size of any charature forst, then we move forward based on its size
	//notice: most of the charaters are 1 byte, so they are increased 1 byte 1 byte,
	//but the above explaintion is for the time that we are using chars that are more
	//then a byte, like chars in other languages

	fmt.Println(strings.Repeat("#", 20))
	str := "hesam"
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c\n", r)
		i += size
	}

	fmt.Println(strings.Repeat("#", 20))
	//2-automatically:
	for _, v := range str {
		fmt.Printf("%c\n", v)
	}
}
