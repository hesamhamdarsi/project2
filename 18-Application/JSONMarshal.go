//we're going to use marshal function from JSOM package to encode a value to JSON
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//notice: in order to marshal function working, we need to define our variables with Upercase. e.g. Name
type person struct {
	Name string
	Age  int
}

func main() {

	person1 := person{
		Name: "Hesam",
		Age:  35,
	}
	//func Marshal(v interface{}) ([]byte, error)
	result, err := json.Marshal(person1)
	if err != nil {
		fmt.Printf("erro: %+v", err)
	}
	//the output is []byte, so you need to convert it to string using one of these methods
	os.Stdout.Write(result)
	fmt.Println("")
	fmt.Println(string(result))

}
