package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome")
}

func webTest() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("server error", err)
	}
}
