package main

import "fmt"

func main() {

	func(id int) {
		fmt.Printf("%+v\n", id)
	}(42)

}
