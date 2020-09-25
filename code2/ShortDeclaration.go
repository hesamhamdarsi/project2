package main

import "fmt"

var y = 50

//zero value for int is 0
var m int

//zero value for string is ""
var n string

var o = "new string"
var p = `this is another 
"new string"`

func main() {
	x := 12
	fmt.Println(x)
	x = 13
	fmt.Println(x)
	y := x + 12
	fmt.Println(y)
	z := "hesam, ali"
	fmt.Println(z)
	foo()
	fmt.Println("variable m before assigning value is=", m)
	fmt.Println("variable n before assigning value is=", n)
	m = 2
	n = "hesam"
	fmt.Println("variable m after assigning value is=", m)
	fmt.Println("variable n after assigning value is=", n)
	fmt.Println("variable p after assigning value is=", p)
	showType()
}

func foo() {
	fmt.Println("value of y=", y)
}
func showType() {
	fmt.Print("type of y is ")
	//format printing
	fmt.Printf("%T\n", y)
	fmt.Print("type of o is ")
	fmt.Printf("%T\n", o)
	//binary, hex, 0x+hex
	fmt.Printf("%b\t%x\t%#x\n", y, y, y)
	//StringPrint. the output would be a string and it should be assign to another value
	x := fmt.Sprintf("%b\t%x\t%#x\n", y, y, y)
	fmt.Println(x)
}
