package main

import "fmt"

func main() {
	slice1 := []string{"Hesam", "Hamed"}
	slice2 := []string{"Majid", "Reza", "Ali"}
	slice4 := []string{"Majid", "ff", "ff"}
	slice3 := [][]string{slice1, slice2, slice4}
	fmt.Println(slice3)
}
