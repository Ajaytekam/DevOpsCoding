package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage ./http-get-json <url>\n")
		return
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("Usage: ./http-get-json <url>\n\nError: Not valid url [%s]", args[1])
		os.Exit(1)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, string(body))
		os.Exit(1)
	}

	var words Words

	err = json.Unmarshal(body, &words)

	fmt.Printf("JSON: Parsed\n\nPage: %s\nWords: %s\n", words.Page, strings.Join(words.Words, ", "))
}
