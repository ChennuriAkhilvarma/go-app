package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, DevOps World!")
}

// #scdsvdvbfdfcvbkmckv gkb bk
func main() {
	http.HandleFunc("/hello", HelloHandler)
	log.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
//cdmflvxmlvbmvS