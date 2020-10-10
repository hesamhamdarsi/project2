package main

import "fmt"

func main() {

	fun1()

	//using argument
	fun2("Hesam")

	//return a result
	result := fun3("book")
	fmt.Println(result)

	//return multiple result
	age := 18
	result1, result2 := fun4("ladugaga", age)
	if result2 == false {
		fmt.Printf("as your age is %d , you are not permited to use this book :'%s'", age, result1)
	} else {
		fmt.Printf("you are ok to use this book\n")
	}
}
func fun1() {
	fmt.Println("the first type")
}
func fun2(name string) {
	fmt.Println("My name is", name)
}

// func <identifier> (argument type) returning type
func fun3(name string) string {
	return fmt.Sprint("that's a greate ", name)
}

//multiple input and output
func fun4(name string, age int) (string, bool) {
	a := name
	b := true
	if age < 16 {
		b = false
	}
	return a, b
}
