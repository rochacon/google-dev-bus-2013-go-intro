package main

import (
    "fmt"
    "sync"
)

var Counter struct {
    sync.RWMutex
    x int
}

func main() {
    wg := &sync.WaitGroup{}
    wg.Add(3)
    go sum(wg)
    go sum(wg)
    go sum(wg)
    wg.Wait()
    fmt.Println("x =", Counter.x)
}

func sum(wg *sync.WaitGroup) {
    for x := 0; x < 10; x++ {
        Counter.Lock()
        Counter.x++
        Counter.Unlock()
    }
    wg.Done()
}
