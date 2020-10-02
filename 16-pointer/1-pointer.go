//pointers are a type
//pointers shows the address of memory
//pointers type starts with *
package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%+v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%+v\n", &a)
	fmt.Printf("%T\n", &a)
	var b *int = &a
	fmt.Printf("%+v\n", b)
	//if you have a variable which is address, you can dereference it using *, this will give you
	//the value which is stored on that address
	fmt.Printf("%+v\n", *b)

}
