//Sort works by a slice of struct
package main

import (
	"fmt"
	"sort"
)

type Cars struct {
	model string
	Speed int
	Plate int
}

//To Sort by Speed
//--------------------------------------------------//
//first we need to make a method for sorting by Speed that will accept slice of anthing(here slice of type Cars)
type Byspeed []Cars

//according to the documentation of Sort sunction, we need to use these 3 function
func (s Byspeed) Len() int           { return len(s) }
func (s Byspeed) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Byspeed) Less(i, j int) bool { return s[i].Speed > s[j].Speed }

//To sort by Model
//--------------------------------------------------//
//the same
type ByModel []Cars

func (m ByModel) Len() int           { return len(m) }
func (m ByModel) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByModel) Less(i, j int) bool { return m[i].model < m[j].model }

func main() {
	//the way we push a bunch of struct in to a slice. []Struct{{Struc1},{Struc2},...,{Strucn}}
	CarList := []Cars{
		{model: "Benz",
			Speed: 300,
			Plate: 230498},
		{model: "Audi",
			Speed: 200,
			Plate: 7849},
		{model: "Megan",
			Speed: 310,
			Plate: 324565},
		{model: "Volvo",
			Speed: 290,
			Plate: 68349},
	}

	fmt.Println(CarList)
	sort.Sort(Byspeed(CarList))
	fmt.Println(CarList)
	sort.Sort(ByModel(CarList))
	fmt.Println(CarList)

}
