package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := "My name is Hesam"
	fmt.Println(s[0])
	fmt.Println(string(s[0]))
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	//converting string to slice of bytes (asci codes)
	ConvertedToByte := []byte(s)
	fmt.Println(ConvertedToByte)
	fmt.Printf("%T\n", ConvertedToByte)
	//printing UTF-8 asci of each character inside the string
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Println(i, v)
	}
	for i, v := range s {
		//print hex
		fmt.Printf("the %d'th value is %#x\n", i, v)
		//bianry
		fmt.Printf("the %d'th value is %b\n", i, v)
		//base-10
		fmt.Printf("the %d'th value is %d\n", i, v)
		fmt.Println("")
	}

	//number to string and wise versa:
	//normally if you want to convert int to string, it will convert to asci code
	//normally you cannot convert float to string
	//but if you use fmt.Sprintf("%d" , int_number)
	//and
	//fmt.Sprintf("%f" , float_number) and put that in a variable, it will turn to string
	var myString = fmt.Sprintf("%d", 222)
	fmt.Printf("%T\n", myString)
	var myString2 = fmt.Sprintf("%d", 22.2222)
	fmt.Printf("%T\n", myString2)

	fmt.Println("")
	//string to int,float,bool,uint:
	s1 := "String"
	fmt.Printf("%T\n", s1)
	var f1, err = strconv.ParseInt(s1, 10, 64) //string, base10 int, 64
	_ = err
	fmt.Printf("%T\n", f1)
	var f2, err2 = strconv.ParseFloat(s1, 64)
	_ = err2
	fmt.Printf("%T\n", f2)

	fmt.Println("")

	//Asci to int and int to asci
	var f3, err3 = strconv.Atoi("-50")
	_ = err3
	fmt.Printf("%T\n", f3)
	fmt.Println(f3)

	var f4 = strconv.Itoa(20)
	fmt.Printf("%T\n", f4)
	fmt.Println(f4)
}
