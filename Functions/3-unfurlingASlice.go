package main

import "fmt"

func main() {

	//use slice as your entry point to the function as variadic parameters
	myslice := []int{1, 2, 3, 4, 5, 6}

	//we need to use ... as we want to say this value is unlimited, otherwise it doesnt accept slice type
	sum, multi := calc(myslice...)

	fmt.Printf("the sum result is %d \n", sum)
	fmt.Printf("the multiplier result is %d \n", multi)
	//////////////////////////////////////////////////////
	//example2:
	//sending null
	sum, multi = calc()
	fmt.Printf("the sum result is %d \n", sum)
	fmt.Printf("the multiplier result is %d \n", multi)

}

func calc(number ...int) (int, int) {
	sum := 0
	multi := 1
	for i, v := range number {
		sum = sum + v
		multi = multi * v
		//why I used i=i? because when we use i in for loop, we have to use it somewhere.
		// otherwise we get error
		i = i
	}
	return sum, multi
}
