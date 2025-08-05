package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, DevOps World!")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)

	// FIX: handle the error returned by ListenAndServe
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
