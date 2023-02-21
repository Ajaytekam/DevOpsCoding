// Example of unidirectional chennal variables  
package main 

import (
    "fmt"
    "time"
)

// Only send data to channel 
func SayHello(c chan<- string) {  
    for i :=0;;i++ {
        c <- "Hello world"
    }
}

// Only receive data to channel  
func printMsg(c <-chan string) {
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


