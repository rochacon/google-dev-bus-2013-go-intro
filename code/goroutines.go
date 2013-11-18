package main

import (
	"fmt"
	"time"
)

func main() {
	go concurrentHello()
	time.Sleep(time.Second)
}

func concurrentHello() {
	fmt.Println("Hello world!")
}
