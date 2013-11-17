package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening on 0.0.0.0:8000")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go!"))
}
