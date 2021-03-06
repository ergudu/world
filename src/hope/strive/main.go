package main

import (
	"net/http"
	"fmt"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w, "hello go")
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
