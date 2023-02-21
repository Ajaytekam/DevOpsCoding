package main

import (
    "fmt"
    "time"
)

func SetMessage(ch chan<- string, msg string) {
    for i:=0;i<5;i++ {
        ch <- msg
    }
}

func main() {
    ch := make(chan string, 10)

    go SetMessage(ch, "Message One")
    go SetMessage(ch, "Message Two")

    // wait for 5 second
    time.Sleep(5*time.Second)
    // close the channel
    close(ch)

    for v := range ch {
        fmt.Println(v)
        time.Sleep(100*time.Millisecond)
    }
}
