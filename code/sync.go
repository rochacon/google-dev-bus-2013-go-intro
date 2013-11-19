// +build OMIT
package main

import (
	"fmt"
	"net/http"
	"sync"
)

// STARTCOUNTER OMIT
var SuccessCounter struct {
	sync.Mutex
	x int
}

// STOPCOUNTER OMIT

// STARTMAIN OMIT
func main() {
	urls := []string{
		"http://google.com",
		"http://golang.org",
		"http://github.com/rochacon",
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(urls))
	for _, url := range urls {
		go checkHealth(wg, url)
	}
	wg.Wait()
	fmt.Println(SuccessCounter.x, "Successfully hitted URLs")
	fmt.Println(len(urls)-SuccessCounter.x, "Failed hits")
}

// STOPMAIN OMIT

// STARTCHECK OMIT
func checkHealth(wg *sync.WaitGroup, url string) {
	r, err := http.Get(url)
	if err != nil {
		return
	}
	if r.StatusCode == http.StatusOK {
		SuccessCounter.Lock()
		SuccessCounter.x++
		SuccessCounter.Unlock()
	}
	wg.Done()
}

// STOPCHECK OMIT
