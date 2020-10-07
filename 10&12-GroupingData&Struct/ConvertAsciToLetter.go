package main

import "fmt"

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
}
