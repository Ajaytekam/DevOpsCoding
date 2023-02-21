package main

import (
    "fmt"
    "time"
)

func AddMsg(ch chan<- string, msg string) {
    for i:=0;i<5;i++ {
        ch <- msg
    }   
    close(ch)
}

func main() {
    ch := make(chan string, 5)   
    go AddMsg(ch, "Hello World")
    time.Sleep(2*time.Second) 
    for msg :=  range ch {
        fmt.Println(msg)
    } 
}
