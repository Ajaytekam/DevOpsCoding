package main 

import (
    "fmt"
    "time" 
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func(){
        for{
            ch1 <- "Message from Func1"
            time.Sleep(1*time.Second)
        }
    }()


    go func(){
        for{
            ch2 <- "Message from Func2"
            time.Sleep(2*time.Second)
        }
    }()

    for { 
        select {
            case msg1 := <- ch1:
                fmt.Println(msg1)
            case msg2 := <- ch2:
                fmt.Println(msg2)
            case <- time.After(1*time.Second):
                fmt.Println("No messages received: Time out")
        }
    }
}

