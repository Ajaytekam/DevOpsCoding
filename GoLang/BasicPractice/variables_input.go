package main

import "fmt"

func main() {
    var userName string
    var age int 
    var gender string 
    fmt.Println("Enter Following Details: ")
    fmt.Printf("Name:> ")
    fmt.Scan(&userName)
    fmt.Printf("Age:> ")
    fmt.Scan(&age)
    fmt.Printf("Gender:> ")
    fmt.Scan(&gender)
    fmt.Println("Name:", userName)
    fmt.Println("Age:", age)
    fmt.Println("Gender:", gender)
}


