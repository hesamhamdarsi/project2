package main

//recursion means calling a function inside itself
//it normally uses for loop
//we recommend using loop instead as it take less memory than recursion
import "fmt"

func main() {

	x := multiplier(5)
	fmt.Println(x)
}
func multiplier(number int) int {
	if number == 0 {
		return 1
	}
	return number * multiplier(number-1)
}
