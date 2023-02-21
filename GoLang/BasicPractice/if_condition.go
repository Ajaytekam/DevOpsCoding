package main

import (
    "fmt"
    "strings"
)

func main() {
    var userName string
    var age int 
    var gender string 
    var preffix string
    fmt.Println("Enter Following Details: ")
    fmt.Printf("Name:> ")
    fmt.Scan(&userName)
    fmt.Printf("Age:> ")
    fmt.Scan(&age)
    fmt.Printf("Gender:> ")
    fmt.Scan(&gender)
    gd := strings.ToLower(gender)
    if (gd == "male") {
        preffix = "Mr."
    } else if (gd == "female"){
        preffix = "Miss"
    } else {
        preffix = ""
    }
    fmt.Printf("Hello %v %v. You are %v years old.\n", preffix, userName, age)
}
