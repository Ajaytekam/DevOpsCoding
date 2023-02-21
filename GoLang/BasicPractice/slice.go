package main

import "fmt"

func main() {
    var nameLists []string
    nameLists = append(nameLists, "John")
    nameLists = append(nameLists, "Snow")
    nameLists = append(nameLists, "Mohan")
    nameLists = append(nameLists, "Pramod")
    //fmt.Println(nameLists)
    for _,j := range nameLists {
        fmt.Println(j)
    }
}
