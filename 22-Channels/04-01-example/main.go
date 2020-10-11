//in this example, we check the stauts code of given urls and we store the body of each succeed URL in a file with it's name
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func checkAndSaveBofy(url string, ch chan string) {
	response, err := http.Get(url)
	if err != nil {
		//as we want to store all values in the channel buffer, we use Sprintf instead
		s := fmt.Sprintf("Error: Site %s is Down!, %v", url, err)
		s += fmt.Sprintf(strings.Repeat("#", 20))
		ch <- s
	} else {
		//close the web session
		defer response.Body.Close()
		s := fmt.Sprintf("%s status code is: %v \n", url, response.StatusCode)
		if response.StatusCode == 200 {
			bodybytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				s += fmt.Sprintf("%v", err)
			}

			file := strings.Split(url, "//")[1]
			file += ".txt"

			s += fmt.Sprintf("Writing response body to %s \n", file)

			err = ioutil.WriteFile(file, bodybytes, 0644)
			if err != nil {
				s += fmt.Sprintf("%v", err)
			}
		}
		s += fmt.Sprintf(strings.Repeat("#", 20))
		ch <- s
	}
}

func main() {
	urls := []string{"https://google.com", "https://yahoo.com", "https://github.com"}
	//we don't need to waitGroup anymore
	//var wg sync.WaitGroup
	ch := make(chan string)
	defer close(ch)
	for _, url := range urls {
		//as we are defining the wait group locally in main function, we need to pass its address to the function, so that it can close it
		go checkAndSaveBofy(url, ch)
		//fmt.Println(strings.Repeat("#", 20))

	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
}
