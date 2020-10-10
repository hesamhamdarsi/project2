package main

import "fmt"

type license struct {
	lisenceID     int
	lisenceDate   string
	level2Lisence bool
}

type person struct {
	name     string
	id       int
	licenses license
}

//no matter if the body is empty. you need to only define person as a deriverInfo() method
func (p person) deriverInfo() {
	fmt.Printf("name: %s\nID: %d\nlisence ID: %d\nlisenceDate: %s\nLevel2: %+v \n", p.name, p.id, p.licenses.lisenceID, p.licenses.lisenceDate, p.licenses.level2Lisence)
}

//no matter if the body is empty. you need to only define license as a deriverInfo() method
func (l license) deriverInfo() {
	fmt.Printf("lISENCE ID: %d\n", l.lisenceID)
}

// we say any value that has the deriverInfo() method, has also the following type
//notice: if the body of interface was emtpy, that means that all values could be belonged to this interface
type driverskill interface {
	deriverInfo()
}

func speeding(d driverskill) {
	//to access to the body of struct, we use (struct).
	//to access to body of any type actually we can use (type), e.e. (string)
	//actually we give the interface as argument to this function
	//the interface is represent our method
	//the method is belonged to the struct and can change it's content.
	//so we can access to the content of the struct through this function
	// notice: both person and type are defined as deriverInfo() method
	switch d.(type) {
	case person:
		fmt.Printf("Driver '%s' with lisence id '%d' has good skillset\n", d.(person).name, d.(person).licenses.lisenceID)
	case license:
		fmt.Printf("Driver with lisence id '%d' \n", d.(license).lisenceID)
	}
}

func main() {

	person1 := person{
		name: "Hesam",
		id:   1234566,
		licenses: license{
			lisenceID:     12,
			lisenceDate:   "1399/5/6",
			level2Lisence: true,
		},
	}

	licenses1 := license{
		lisenceID:     13,
		lisenceDate:   "1399/5/6",
		level2Lisence: false,
	}
	//no matther if you call this method or not, anyway you can use speeding() for both structs
	person1.deriverInfo()
	licenses1.deriverInfo()
	speeding(person1)
	speeding(licenses1)

}
