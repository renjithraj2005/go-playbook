package main

import (
	"fmt"
	"io/ioutil"
	"log" //Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output.
	"net/http"
	"time"
)

func fetch() {
	url := "https://www.google.co.in/"
	t := time.Now()
	ping(url)

	fmt.Println()
	fmt.Print(time.Now().Sub(t))
}

// Golang http documentation : https://golang.org/pkg/net/http/
func ping(url string) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1573.2 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}

	if response.StatusCode == 200 {
		fmt.Println("success")
		respstrem, err := ioutil.ReadAll(response.Body)
		response.Body.Close()

		if err != nil {
			log.Fatal(err.Error())
		}
		html := string(respstrem)
		fmt.Println(html)
	}
}
