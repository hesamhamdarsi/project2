//func CompareHashAndPassword(hashedPassword, password []byte) error
//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
//  >> crypt.GenerateFromPassword([]byte(password), bcrypt.cost) (output:2 variables)
package main

import (
	"fmt"
	//for these packages that are exprimental (x) and they are not in standard liberary yet
	//you can use "go get golang.org/x/crypto/bcrypt" in your system to pull the package to your system
	"golang.org/x/crypto/bcrypt"
)

func main() {

	MyPassword := "hesam123"
	pass, err := bcrypt.GenerateFromPassword([]byte(MyPassword), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(MyPassword)
	fmt.Println(pass)

	err = bcrypt.CompareHashAndPassword(pass, []byte(MyPassword))
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("Suucessfully logged in")
}
