package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func main() {

	http.HandleFunc("/", index)

	http.ListenAndServe("0.0.0.0:8000", nil)
}
