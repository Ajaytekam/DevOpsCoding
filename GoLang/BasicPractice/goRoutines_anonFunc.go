package main

import (
    "fmt"
    "time"
)

func main() {

    go func() {
        for i:=0;i<6;i++{
            fmt.Println("01: Hello world")
            time.Sleep(time.Second*1)    
        }
    }()    


    go func() {
        for i:=0;i<3;i++{
            fmt.Println("02: Hello Test")
            time.Sleep(time.Second*2)    
        }
    }()    

    var str string
    fmt.Scanln(&str)
}
