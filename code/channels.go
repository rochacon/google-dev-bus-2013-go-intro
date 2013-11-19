// +build OMIT
package main

import (
	"fmt"
	"net/http"
)

// START0 OMIT
func main() {
	urls := []string{
		"http://google.com",
		"http://golang.org",
		"http://github.com/rochacon",
	}

	isSuccess := make(chan bool, len(urls))

	for _, url := range urls {
		go checkHealth(url, isSuccess)
	}

	success := 0
	for _, _ = range urls {
		if <-isSuccess {
			success++
		}
	}

	fmt.Println(success, "Successfully hitted URLs")
	fmt.Println(len(urls)-success, "Failed hits")
}

// STOP0 OMIT

// START1 OMIT
func checkHealth(url string, succeeded chan<- bool) {
	r, err := http.Get(url)
	if err != nil {
		succeeded <- false
		return
	}
	succeeded <- r.StatusCode == http.StatusOK
}

// STOP1 OMIT
