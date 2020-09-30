package main

import "fmt"

func main() {

	//struct is like a class. we make it and then we can make object of that class
	// we define a type for that. since now, this struct would be a special type for us
	//we call that also aggregate data structure, composite data structure, complex data type
	// type <Name> <Data structure type>
	//like the other types you can define it above main func but only if it needs to be present out of main
	type personalInfo struct {
		FirstName string
		LastName  string
		ID        int
	}
	//objectA
	PersonA := personalInfo{
		FirstName: "Hesam",
		LastName:  "Hamdarsi",
		ID:        82094578,
	}
	//objectB
	PersonB := personalInfo{
		FirstName: "Hesam",
		LastName:  "Hamdarsi",
		ID:        82094578,
	}
	fmt.Println(PersonA)
	fmt.Println(PersonB)
	fmt.Println(PersonA.FirstName, PersonA.LastName, PersonA.ID)
	fmt.Println(PersonB.FirstName, PersonB.LastName, PersonB.ID)

	//Embeded struct (nested struct)
	type position struct {
		pi        personalInfo
		positions string
	}

	//notice: personalInfo : personalInfo...
	//as you see the feild which is defined in position
	PosPersonA := position{
		pi: personalInfo{
			FirstName: "Hesam",
			LastName:  "Hamdarsi",
			ID:        82094578,
		},
		positions: "Manager",
	}
	PosPersonB := position{
		pi: personalInfo{
			FirstName: "majid",
			LastName:  "Majidi",
			ID:        12121212,
		},
		positions: "senior engineer",
	}

	fmt.Println(PosPersonA)
	fmt.Println(PosPersonB)
	fmt.Println(PosPersonA.pi.FirstName, "-", PosPersonA.pi.LastName, "-", PosPersonA.pi.ID, "-", PosPersonA.positions)
	fmt.Println(PosPersonB.pi.FirstName, "-", PosPersonB.pi.LastName, "-", PosPersonB.pi.ID, "-", PosPersonB.positions)
}
