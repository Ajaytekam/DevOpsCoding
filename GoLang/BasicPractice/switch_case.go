package main

import (
    "fmt"
    "strings"
)

func main() {
    var ch string

    fmt.Printf("Enter your choice [Y|N]: ")
    fmt.Scan(&ch)

    switch strings.ToLower(ch) {
        case "n":
            fmt.Println("You selected No")
        case "y":
            fmt.Println("You selected Yes")
        default:
            fmt.Println("Did not understand your choice!!")
    }
}  
