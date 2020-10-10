//like functions, we have two types: pointer based and value based
//if we want to chenge the value of variables in a method, we need to use a pointer receiver

package main

import (
	"fmt"
)

// Author structure
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a *author) show() {

	(*a).name = "Hesam"
}

// Main function
func main() {

	res := author{
		name:      "Sona",
		branch:    "CSE",
		particles: 203,
		salary:    34000,
	}

	(&res).show()
	//notice: golang allows us to use the variable instead of &var. it will correct the entry itseld
	//so the below would be ok as well
	//res.show()
	fmt.Println(res)

}
