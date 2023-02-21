package main

import (
    "fmt"
)

func SetMsg(ch1 chan<- string) {
// defining ch1 to send data
    ch1 <- "Hello world"
}

func GetMsg(ch1 <-chan string, ch2 chan<- string) {
// defining ch1 to receive data and ch2 to send data
    msg := <- ch1 
    ch2 <- msg
}

func main() {
    // declaring chennals with buffered capacity of 1
    ch1 := make(chan string, 1)
    ch2 := make(chan string, 1)
    
    SetMsg(ch1)
    GetMsg(ch1,ch2)
    fmt.Println(<-ch2)
}
