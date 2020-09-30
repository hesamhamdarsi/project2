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
		fmt.Printf("%c\n", rune(i))

	}
}
