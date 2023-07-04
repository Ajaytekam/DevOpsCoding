package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Users struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

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
	// fmt.Printf("%s\n",body)

	var userDetails []Users

	// Method 01
	// decode the json file into slice of users struct
	// decoder := json.NewDecoder(bytes.NewReader(body))
	// if err = decoder.Decode(&userDetails); err != nil {
	//     fmt.Println(err)
	//     return
	// }

	// Method 02
	// json.Unmarshall() parse []byte array json data fields to the corresponding fields in the Go struct.
	err = json.Unmarshal(body, &userDetails)
	if err != nil {
		log.Fatal(err)
	}

	// Print the parsed data
	for _, user := range userDetails {
		fmt.Printf("ID: %d\nName: %s\nEmail: %s\nGender: %s\nStatus: %s\n\n", user.Id, user.Name, user.Gender, user.Email, user.Status)
	}
}
