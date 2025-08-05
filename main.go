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
	http.ListenAndServe(":8080", nil)
}
//vbfdf,;fd
//vdsfgrfdsrfvr
//sbvdfdfdbfe
//bfaesgfesb
//dbfdfesfwaefeds//