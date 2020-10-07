package main

import "fmt"

func main() {
	//array
	var myArray [5]int
	myArray[3] = 2
	fmt.Println(myArray)
	fmt.Println(len(myArray))
	fmt.Println(myArray[4])
	//first method of creating array
	for i, v := range myArray {
		fmt.Printf("%+v-%+v\n", i, v)
	}

	//second method of creating array
	myArray2 := [5]int{1, 2, 3, 4, 5}
	for i, v := range myArray2 {
		fmt.Printf("%+v-%+v\n", i, v)
	}

	//composite litteral: to group together values of the same type
	//composite litteral construct values for structs, arraye and maps
	//this one is slice. slices are like array but you wont get them fixed size and they are enabled
	//to resize
	//https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html
	myComposite := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(myComposite)
	fmt.Println(len(myComposite))
	fmt.Println(myComposite[2])

	for i, v := range myComposite {
		fmt.Printf("%+v-%+v\n", i, v)
	}

	//slice of slice
	fmt.Println(myComposite[:])
	fmt.Println(myComposite[2:])
	fmt.Println(myComposite[:4])
	fmt.Println(myComposite[1:3])

	//append to slice
	//placeholder
	myComposite = append(myComposite, 7, 8, 9)
	fmt.Println(myComposite[:])

	//append slice to slice
	//... is mandetory to ignore the error
	myNewComposite := []int{20, 30, 40}
	myComposite = append(myComposite, myNewComposite...)
	fmt.Println(myComposite[:])

	//deleting from slice, for instace we want to remove 3,4
	myComposite = append(myComposite[:2], myComposite[4:]...)
	fmt.Println(myComposite[:])

}
