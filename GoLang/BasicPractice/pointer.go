package main

import "fmt"

func main() {
    var ptr *int
    number := 10
    ptr = &number
    fmt.Println("Value: ", *ptr)
    fmt.Println("Address: ", ptr)
    fmt.Printf("Pointer-Type: %T\n", ptr)
}
