//pointers are a type
//pointers shows the address of memory
//pointers type starts with *
//pointer *int means address where an int is stored
package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

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

	//#########################################
	fmt.Println(strings.Repeat("#", 20))
	//#########################################

	//if you want to change the value of a variable using another function, you need to pass
	//the address of that variable instead of the variable itself

	x := 10
	fmt.Println("x: ", x)
	changeVar(x)
	fmt.Println("x after call the function without using pointer: ", x)
	changeVarPointer(&x)
	fmt.Println("x after call the function without using pointer: ", x)

	//#########################################
	fmt.Println(strings.Repeat("#", 20))
	//#########################################

	//if you want to change the value of a struct using another function, you need to pass
	//the address of that struct instead of the variable itself

	st := person{"hesam", 35}
	fmt.Println("x: ", st)
	changeStruct(st)
	fmt.Println("x after call the function without using pointer: ", st)
	changeStructPointer(&st)
	fmt.Println("x after call the function without using pointer: ", st)

	//#########################################
	fmt.Println(strings.Repeat("#", 20))
	//#########################################

	//if you want to change the value of a slice or Map using another function, you Don't need
	//to send the address, because they actually don't store the actual data but refer to a memory address

	prices := []int{10, 20, 30}
	fmt.Println("prices: ", prices)

	changeSlice(prices)
	fmt.Println("Prices after call the function without using pointer: ", prices)

	UserAge := map[string]int{"a": 10, "b": 20}
	fmt.Println("UserAge: ", UserAge)
	changeMap(UserAge)
	fmt.Println("Prices after call the function without using pointer: ", UserAge)

}
func changeVar(x int) {
	x = 20
}
func changeVarPointer(x *int) {
	*x = 20
}

func changeStruct(st person) {
	st.age = 20
}
func changeStructPointer(pt *person) {
	(*pt).age = 20
	//OR
	//pt.age = 20
	//struct can change the value using the pointer itself or the value

}

func changeSlice(s []int) {
	for i := range s {
		s[i] *= 2
	}

}

func changeMap(m map[string]int) {
	m["a"] = 20
	m["b"] = 40
}
