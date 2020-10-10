package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	//creating File
	//notice: you can create the variables before using them, or you can create and declarate them
	//simultanously like below
	myFile, err := os.Create("myFile.txt")
	if err != nil {
		//fmt.Println(err)
		//os.Exit(1)

		log.Fatal(err)
	}
	//if you wanted to first define the variables then use them you should:
	//var myFile *os.NewFile       >> when you create a file, its a pointer type acually
	//var err error                >> error type

	//notice: you can use the path as well, but consider the linux, windows, etc path are different
	//e.g. os.Create("/etc/a.txt")

	//##############################################################
	fmt.Println(strings.Repeat("#", 20))
	//##############################################################

	//truncating a file (limiting the size of the file and remove the extra data or make it totaly empty)
	err = os.Truncate("myFile4.txt", 50)
	if err != nil {
		fmt.Println(err)
	}
	myFile.Close()
	//to check
	info, err := os.Stat("myFile4.txt")
	fmt.Println("size in byte= ", info.Size())

	//##############################################################
	fmt.Println(strings.Repeat("#", 20))
	//##############################################################

	//open a file method 1
	myFile, err = os.Open("myFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	myFile.Close()

	//open a file method 2
	//open the file-append to the file (O_append)-set permission
	//you can use other flages like os.O_RDONLY, etc
	myFile, err = os.OpenFile("myFile.txt", os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	myFile.Close()

	//open a file method 3
	//open the file-if not exist, create it and then append to the file (O_append)-set permission
	myFile, err = os.OpenFile("myFile2.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	myFile.Close()

	//##############################################################
	fmt.Println(strings.Repeat("#", 20))
	//##############################################################

	//getting file info
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("myFile.txt")
	p := fmt.Println

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}
	p("File Size(byte): ", fileInfo.Size())
	p("File Name: ", fileInfo.Name())
	p("Last Modify: ", fileInfo.ModTime())
	p("Is Directory? ", fileInfo.IsDir())
	p("Permissions: ", fileInfo.Mode())

	//##############################################################
	fmt.Println(strings.Repeat("#", 20))
	//##############################################################

	//Rename a file
	oldPath := "myFile.txt"
	newPath := "myFilesss.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println(err)
	}

	//##############################################################
	fmt.Println(strings.Repeat("#", 20))
	//##############################################################

	//Remove a file
	err = os.Remove("myFile3.txt")
	if err != nil {
		fmt.Println(err)
	}
	//err = os.Remove("myFilesss.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}

}
