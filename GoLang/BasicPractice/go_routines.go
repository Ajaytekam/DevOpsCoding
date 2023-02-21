package main 

import (
    "fmt"
    "time"
)

func printT(tNo int, t int, delay int) {
    for i:=0;i<t;i++ {
        fmt.Println("ThreadNo:", tNo, "| Delay:", delay, "seconds | Counter:",i)
        time.Sleep(time.Second * time.Duration(delay))
    }
}

func main() {    
    go printT(1, 5, 3)    
    go printT(2, 6, 2)    
    go printT(3, 7, 1)    
    var str string
    fmt.Scanln(&str)
}
