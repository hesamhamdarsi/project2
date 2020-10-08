//this function will do both test and documentation together
//this will make a document form the package file (calculate package here)
//output:
//result
//the above strings, will bring the result of the function we are calling from package
//you can check the created document with running the following command:
//godoc -http=:8088 (or any port you want)
//then you can open that in your bowser and go to package section and search for the package you just created
package calculate

import "fmt"

//the other types of Example:
//https://golang.org/pkg/testing/#Examples
func ExampleSum() {
	fmt.Println(Sum(1, 2, 3))
	//Output:
	//6
}
