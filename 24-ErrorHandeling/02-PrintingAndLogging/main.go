//fmt.println()
//log.Println(): you can spesify the output file,
//               and it give you the date, timestamps
//log.Fatalln(): after error the progrma will exit immidiately
//               actually it runs the os.exit()
//               show error status
//log.Panicln()  >> after that, all deffered functions would be run
//panic()

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//////////////////////////////////////////
	//fmt.println
	_, err := os.Open("nofile.txt")
	if err != nil {
		fmt.Println("err happened: ", err)
	}
	//////////////////////////////////////////
	//log.println
	logfile, err := os.Create("logfile.txt")
	if err != nil {
		fmt.Println("err happened: ", err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)
	f2, err := os.Open("nofile.txt")
	if err != nil {
		log.Println("err happened: ", err)
	}
	defer f2.Close()
	//////////////////////////////////////////
	//log.Fatalln
	//	_, err = os.Open("nofile.txt")
	//	if err != nil {
	//log.Fatalln(err)
	//	}
	//////////////////////////////////////////
	//////////////////////////////////////////
	//log.panicln
	_, err = os.Open("nofile.txt")
	if err != nil {
		log.Panicln(err)
		//panic(err)
	}
	//////////////////////////////////////////
}
