//Golang don't have classes, but you an define methods on predefined types
//Named Types can have functioned attached to them which we call it method
//methods can work on both data and behavious
//type ,ay have methodset associated with it which enhanced the type with extra behavior
//the major difference between a method and a function is that the method belongs to a type but
//function belongs to a a pckage

package main

import (
	"fmt"
	"time"
)

// Author structure
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

type names []string

func (n names) prints() {
	for _, v := range n {
		fmt.Println(v)
	}
}

// Method with a receiver
// of author type
// func(reciver_name Type) method_name(parameter_list)(return_type)
// we pass the receiver as an argument and change it's properties
func (a author) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// Main function
func main() {

	//example1:
	//using the method from a apckage in standard liberary
	//here the type of constant would be time.duration. actually time.duration is a type made by golang
	//and through that we've made a named type of value called "day".
	const day = 24 * time.Hour
	fmt.Println("Example1")
	fmt.Printf("the type of day is %T\n", day)
	//now we use the seconds() method of time.duration type. this is a method also known as receiver function
	seconds := day.Seconds()
	fmt.Printf("%+v\n", seconds)
	//so the named values can have functions, also called methods or receiver functions attached to them

	fmt.Println("###############################################")

	//example2: creating a method
	//we create a method of slice of strings
	//then we create a method for this new type
	//then any value of this type gets access to the created method
	//type names []string
	//method: func (receiver) name(){}
	//the receiver is a copy of the type we are working on
	//so here the underline type of names is slice of strings. so we can do anything with it that we
	//normally do with slice of strings
	//notice: the method and type should be defined outside of the main() function body

	freinds := names{"hesam", "ali"}
	fmt.Println("Example2")
	freinds.prints()
	//another way to call the method:
	names.prints(freinds)

	fmt.Println("###############################################")

	//example3: method and struct
	// Initializing the values
	// of the author structure
	res := author{
		name:      "Sona",
		branch:    "CSE",
		particles: 203,
		salary:    34000,
	}

	// Calling the method
	res.show()

	fmt.Println("###############################################")

	//using function for the same target
	//need each individual parameter to be passed to function as argument
	// but using method, it wasn't need as we pass the receiver as an argument and change it's properties
	funcz(res.name, res.branch, res.particles, res.salary)
}
func funcz(a string, b string, c int, d int) {
	fmt.Println("Author's Name: ", a)
	fmt.Println("Branch Name: ", b)
	fmt.Println("Author's Name: ", c)
	fmt.Println("Branch Name: ", d)
}
