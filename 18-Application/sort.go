package main

import (
	"fmt"
	"sort"
)

func main() {

	Myint := []int{19, -3, 2, 5, 7, 2, 0, 4, 14, 14, 56, 111, 23, 8}
	MyString := []string{"Hesam-", "my name-", "is-", "I-", "Am-", "33years old-"}

	fmt.Println(MyString)
	sort.Strings(MyString)
	fmt.Println(MyString)

	fmt.Println(Myint)
	sort.Ints(Myint)
	fmt.Println(Myint)
}
