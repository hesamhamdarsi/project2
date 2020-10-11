//in this example, we check the stauts code of given urls and we store the body of each succeed URL in a file with it's name
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func checkAndSaveBofy(url string, wg *sync.WaitGroup) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println("Site Is Down")
	} else {
		defer response.Body.Close()
		fmt.Printf("%s status code is: %v \n", url, response.StatusCode)
		if response.StatusCode == 200 {
			bodybytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			file := strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("Writing response body to %s \n", file)

			err = ioutil.WriteFile(file, bodybytes, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println(strings.Repeat("#", 20))
	wg.Done()
}

func main() {
	urls := []string{"https://google.com", "https://yahoo.com", "https://github.com"}
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		//as we are defining the wait group locally in main function, we need to pass its address to the function, so that it can close it
		go checkAndSaveBofy(url, &wg)
	}
	wg.Wait()
}
