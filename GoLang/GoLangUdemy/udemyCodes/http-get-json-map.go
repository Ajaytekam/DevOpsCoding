package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

// Defining Interface 
type Response interface {
	GetResponse() string
}

// Defining interface method for Words object
func (w Words) GetResponse() string {
	return fmt.Sprintf("Words: %s", strings.Join(w.Words, ", "))	
}

// Defining interface method for Occurrence object
func (o Occurrence) GetResponse() string {
	words := []string{}
	for word, occurrence := range o.Words {
		words = append(words, fmt.Sprintf("%s (%d)", word, occurrence))
	}
	return fmt.Sprintf("Words: %s", strings.Join(words, ", "))
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage ./http-get-json <url>\n")
		return
	}

	res, err := doRequest(args[1])
	if err != nil {
		fmt.Printf("Error :%s\n", err)
		os.Exit(1)
	}

	if res == nil {
		fmt.Printf("No response\n")
		os.Exit(1)
	}

	fmt.Printf("Response: %s\n", res.GetResponse())
}

func doRequest(requestURL string) (Response, error) {
	
	if _, err := url.ParseRequestURI(requestURL); err != nil {
		fmt.Printf("Usage: ./http-get-json <url>\n\nError: Not valid url [%s]", requestURL)
		os.Exit(1)
	}

	response, err := http.Get(requestURL)

	if err != nil {
		return nil, fmt.Errorf("Get error: %s", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("ReadAll error: %s", err)	
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, string(body))
	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s", err)
	}

	switch page.Name {
	case "words":
		var words Words
		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, fmt.Errorf("Unmarshal error: %s", err)
		}
		return words, nil

	case "occurrence":
		var occurrence Occurrence
		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			return nil, fmt.Errorf("Unmarshal error: %s", err)
		}
		return occurrence, nil
	}
	return nil, nil
}

// http://192.168.1.14:8080/words?input=words5 
// http://192.168.1.14:8080/occurrence
