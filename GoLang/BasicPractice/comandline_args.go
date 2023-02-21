package main

import (
    "fmt"
    "os"
)

func main() {
    Length := len(os.Args)
    for i:=0;i<Length;i++ {
        fmt.Println(i+1,"Argument:", os.Args[i])
    }
}
