package main

import "fmt"

func main() {
    myFunc := func(x int) int {
        return x*x*x
    }

    num1 := myFunc(10)
    num2 := myFunc(2)
    num3 := myFunc(3)

    fmt.Println("Number1: ", num1)
    fmt.Println("Number2: ", num2)
    fmt.Println("Number3: ", num3)
}
