package main 

import (
    "fmt"
    "time"
)

func SayHello(c chan string) {
    for i :=0;;i++ {
        c <- "Hello world"
    }
}

func printMsg(c chan string) {
    for {
        msg := <- c 
        fmt.Println(msg)
        time.Sleep(time.Second*1)
    }
}

func main() {
    var c chan string = make(chan string) 
    go SayHello(c)
    go printMsg(c)   

    var str string
    fmt.Scanln(&str)
}


