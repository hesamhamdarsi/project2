package main

import (
	"encoding/json"
	"fmt"
)

// Animals ...
//you need to create a slice for your JSON file
//this slice is a set of structs. so we need to first create a struct that has the same fileds as JSON keys
type Animals struct {
	Name  string
	Order string
}

func main() {

	var JSAnimla = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	//slice of created struct
	var animals []Animals
	//func Unmarshal(data []byte, v interface{}) error
	err := json.Unmarshal(JSAnimla, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)
	for _, v := range animals {
		fmt.Println(v.Name, "\t", v.Order)
	}
}
