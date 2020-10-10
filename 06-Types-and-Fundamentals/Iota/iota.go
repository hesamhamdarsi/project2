//IOTA increase the consts that are in the same () gradually

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
	_ = iota
	//this means 1(1*10 times shift to left) = 1(0000000000)= 10000000000
	//so a siota is 1 here it will be 10000000000
	kb = 1 << (iota * 10)
	//as iota inceased here, mb=1(20 times shift)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Println(kb)
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
	//when in the output you see this:
	//  1024		10000000000
	//this means if we convert 1024 which is decimal to binary it would be 10000000000 (2^10+0+0+...+0)

}

//oter examples:
/*
const (
	a = 1-(iota*7)
	b
	c
	d
)
*/
