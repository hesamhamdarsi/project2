package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	//chack if any word is in the string
	result := strings.Contains("I am hesam", "am")
	p(result)

	///////////////////////////////////////////
	fmt.Println(strings.Repeat("#", 20))
	///////////////////////////////////////////

	//chack if any char is in the string
	result = strings.ContainsAny("I am hesam", "xys")
	p(result)

	///////////////////////////////////////////
	fmt.Println(strings.Repeat("#", 20))
	///////////////////////////////////////////

	//chack the number of a char/word in string
	result1 := strings.Count("I am hesam", "a")
	p(result1)
	result1 = strings.Count("I am hesam", "hesam")
	p(result1)

	///////////////////////////////////////////
	fmt.Println(strings.Repeat("#", 20))
	///////////////////////////////////////////

	//Make upper and Lower case
	p(strings.ToUpper("my naMe IS hESam"))
	p(strings.ToLower("my naMe IS hESam"))

	///////////////////////////////////////////
	fmt.Println(strings.Repeat("#", 20))
	///////////////////////////////////////////

	//chach two string without considering upper case and lower case
	p(strings.EqualFold("HESAM", "hesam"))

	///////////////////////////////////////////
	fmt.Println(strings.Repeat("#", 20))
	///////////////////////////////////////////

	//change char/word in the string
	//the first two
	result3 := strings.Replace("192.168.1.1", ".", ":", 2)
	p(result3)
	//all
	result3 = strings.Replace("192.168.1.1", ".", ":", -1)
	p(result3)
	//all
	result3 = strings.ReplaceAll("192.168.1.1", ".", ":")
	p(result3)
	//all words
	result3 = strings.ReplaceAll("my name is my cat", "my", "his")
	p(result3)

	///////////////////////////////////////////
	fmt.Println(strings.Repeat("#", 20))
	///////////////////////////////////////////

	//spliting: it return slice of characters
	result4 := strings.Split("a,b,c", ",")
	p(result4)
	result4 = strings.Split("my name is hesam", "")
	p(result4)

	///////////////////////////////////////////
	fmt.Println(strings.Repeat("#", 20))
	///////////////////////////////////////////

	//joing strings to strings with adding seperator
	s := []string{"My", "Name", "Hesam"}
	result3 = strings.Join(s, "-")
	p(result3)
	result3 = strings.Join(s, " ")
	p(result3)

	///////////////////////////////////////////
	fmt.Println(strings.Repeat("#", 20))
	///////////////////////////////////////////

	//spliting: with new line and white spaces
	result4 = strings.Fields("My name is Mr \n Hesam hamdarsi")
	p(result4)

	///////////////////////////////////////////
	fmt.Println(strings.Repeat("#", 20))
	///////////////////////////////////////////

	//trim: remove all leading and trailing white spaces
	result3 = strings.TrimSpace("\t My name is Mr Hesam hamdarsi    \n")
	p(result3)
	//trim: remove all leading and trailing special characters
	result3 = strings.Trim("...name is Mr Hesam hamdarsi!!??", ".!?")
	p(result3)

}
