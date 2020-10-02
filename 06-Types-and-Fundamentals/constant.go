package main

import "fmt"

//A constant is a data item whose value cannot change during the program’s execution.
//Thus, as its name implies – the value is constant.
const (
	z int = 12
	y     = "my name is hesam"
)
const (
	k = iota
	f = iota
	g
	h = iota
)

func main() {

	const a = 22
	const b = 11.2234
	const c = "Hesam, Ahmed"

	fmt.Printf("%T\t%+v\n", a, a)
	fmt.Printf("%T\t%+v\n", b, b)
	fmt.Printf("%T\t%+v\n", c, c)
	fmt.Printf("%T\t%+v\n", z, z)
	fmt.Printf("%T\t%+v\n\n", y, y)
	//iota is a type of constant which is int and increament automatically
	fmt.Println("IOTA")
	fmt.Println("-------------------")
	fmt.Printf("%T\t%+v\n", k, k)
	fmt.Printf("%T\t%+v\n", f, f)
	fmt.Printf("%T\t%+v\n", g, g)
	fmt.Printf("%T\t%+v\n", h, h)

}
