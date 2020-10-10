package main

import "fmt"

func main() {

	//get unlimited argument of a type and return as much as result you want
	sum, multi := sum(1, 2, 3, 4, 5, 6)

	fmt.Printf("the sum result is %d \n", sum)
	fmt.Printf("the multiplier result is %d \n", multi)

}

func sum(number ...int) (int, int) {
	sum := 0
	multi := 1
	for i, v := range number {
		sum = sum + v
		multi = multi * v
		//why I used i=i? because when we use i in for loop, we have to use it somewhere.
		//or better way is >> for _, v := range number. using _ means dont need to use it
		// otherwise we get error
		i = i
	}
	return sum, multi
}
