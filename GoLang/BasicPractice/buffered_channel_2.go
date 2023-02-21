package main 

import (
    "fmt"
)

func main() {
    // channel with capacity 2 
    ch := make(chan string, 2)
    
    ch <- "Message 1"
    ch <- "Message 2"

    // receiving/priting both messages
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
