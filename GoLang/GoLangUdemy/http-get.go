package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		return
	}

	// url to test : https://gorest.co.in/public/v2/users
	// parse the given url
	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format: %s\n", err)
		os.Exit(1)
	}

	// make get request
	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}

	// close the response
	defer response.Body.Close()

	// read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print the status code and response body
	// fmt.Printf("HTTP Status Code: %d\nBODY:===========\n%s\n", response.StatusCode, body)
	fmt.Printf("%s\n", body)
}
