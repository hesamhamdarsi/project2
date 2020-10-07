//where it is useful?
//through this, we can use "go func()" and so we can use another CPU to process in paralel if we have more than one

package main

import "fmt"

func main() {

	func(id int) {
		fmt.Printf("%+v\n", id)
	}(42)

	x := func(id int) int {
		fmt.Printf("%+v\n", id)
		return id + 1
	}(42)
	fmt.Printf("%+v\n", x)

}
