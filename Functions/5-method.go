package main

import "fmt"

// Author structure
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

// Method with a receiver
// of author type
// func(reciver_name Type) method_name(parameter_list)(return_type)
// using method, it wasn't need as we pass the receiver as an argument and change it's properties
func (a author) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// Main function
func main() {

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
