package main

import "fmt"

const (
	a = iota
	b = iota
	c
	d = iota
)

const (
	//the first line means ignore the first value of iota which is 0
	//the _ means no where, so _=iota means _=0. that means assign the first iota which is 0 to nowhere
	//so the next value of iota would be 1
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	//iota is a type of constant which is int and increament automatically
	fmt.Println("IOTA")
	fmt.Println("-------------------")
	fmt.Printf("%T\t%+v\n", a, a)
	fmt.Printf("%T\t%+v\n", b, b)
	fmt.Printf("%T\t%+v\n", c, c)
	fmt.Printf("%T\t%+v\n\n", d, d)
	fmt.Println("Shift")
	fmt.Println("-------------------")
	fmt.Printf("%+v\t\t%+b\n", kb, kb)
	fmt.Printf("%+v\t\t%+b\n", mb, mb)
	fmt.Printf("%+v\t%+b\n", gb, gb)

}
