//labels work with "breake", "continue" and "go to"

package main

import (
	"fmt"
)

func main() {
	array1 := [3]string{"hesam", "ali", "hasan"}
	array2 := [2]string{"mohsen", "hesam"}

outer:
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if v2 == v1 {
				fmt.Println("we found a match which is ", v2)
				//here we will be out of the both loop. if we were using only "break"
				//then we were out of the first loop only
				break outer
			}
		}
	}

}
