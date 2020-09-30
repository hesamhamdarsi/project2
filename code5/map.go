package main

import "fmt"

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
	myvalue, mystatus := myMap["Mohammadreza"]
	fmt.Println(myvalue)
	fmt.Println(mystatus)

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
	delete(myMap, "hesam")
	fmt.Println(myMap)

	//again, to check if the value exist, we can use if statement
	//before
	delete(myMap, "hesam")
	fmt.Println(myMap)
	//after
	if myvalue, mystatus := myMap["Majid"]; mystatus {
		fmt.Println(myvalue)
		delete(myMap, "Majid")
	}

}
