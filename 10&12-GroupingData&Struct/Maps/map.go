package main

import (
	"fmt"
	"strings"
)

func main() {

	//Map is key value store. [key]value
	myMap := map[string]int{
		"hesam": 82094578,
		"Ali":   001234455,
	}
	fmt.Println(myMap)
	fmt.Println(myMap["hesam"])

	//but this will also print value 0 for any key which is not in our store
	fmt.Println(myMap["Mohammadreza"])

	//to prevent that, we need to check if the value is exist or not
	//we call this process "comma ok idiom"
	v, ok := myMap["Mohammadreza"]
	fmt.Println(v)
	fmt.Println(ok)

	//we saw in the output that value was 0 but the ok status was false, so we can use an "if" for that:
	if myvalue, mystatus := myMap["Mohammadreza"]; mystatus {
		fmt.Println(myvalue)
	}

	//add new key/value to the exsisting MAP
	myMap["Ahmed"] = 448398920
	fmt.Println(myMap)

	for k, v := range myMap {
		fmt.Printf("National ID of %+v is:\t %+v\n", k, v)
	}

	//deleting from Map
	myvalue, mystatus := myMap["hesam"]
	_ = myvalue
	fmt.Println(mystatus)
	delete(myMap, "hesam")
	fmt.Println(myMap)

	//again, to check if the value exist, we can use if statement
	//before
	delete(myMap, "hesam")
	fmt.Println(myMap)
	myvalue, mystatus = myMap["hesam"]
	fmt.Println(mystatus)
	//after
	if myvalue, mystatus := myMap["Majid"]; mystatus {
		fmt.Println(myvalue)
		delete(myMap, "Majid")
	}

	fmt.Println(strings.Repeat("#", 20))
	//Notice: you can use an array as the key in Map, but you cant use slice
	//because slice size is not declarated
	//	var m1 map[[5]int]string //valid
	//	var m1 map[[]int]string  //invalid

	//Notice:
	//you can not input value in Map when it is not initialize
	// myMap := map[string]int          >> not initializd Map
	// myMap := map[string]int{}        >> initialized Map
	// myMap := make(map[string]int)    >> initialized Map

	fmt.Println(strings.Repeat("#", 20))
	//Notice:
	//Maps can not compare directly, you can only compare them with nil
	//but if the key and values are string in both Maps. we can compare them using "Sprintf"
	//it is like we create 2 strings and compare them

	a := map[string]string{"Name1": "hesam", "Name2": "ali"}
	b := map[string]string{"Name1": "ali", "Name2": "hesam"}

	x := fmt.Sprintf("%s", a)
	y := fmt.Sprintf("%s", b)

	if x == y {
		fmt.Println("equal")
	} else {
		fmt.Println("NOT equal")
	}

	fmt.Println(strings.Repeat("#", 20))
	//duplicate, copy and Map in memory
	//map is indicating the map header in the memory
	//map do not keep the values, the values are saved in the memory
	//when you make a copy of a Map, if you change any otem in each of them, the other will be
	//changed as well, because it refrences the same data structure in Memory, because both of them
	//have the same Mac header

	//so what we can do?
	//we should initial an empty Map and using for loop to copy the values of the existing Map

}
