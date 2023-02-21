package main

import "fmt"

func main() {

    var arr1 [5]string
    var arr2 = [5]int{10, 20, 30, 40, 50}
    
    fmt.Println("Capacity of arr1:", cap(arr1))
    fmt.Println("Length of arr2:", len(arr2))

    fmt.Println()

    fmt.Println("First Array Example: ")
    fmt.Println("Array Values")
    for i,_ := range arr2 {
        fmt.Println(arr2[i])
    }

    fmt.Println()

    fmt.Println("Second Array Example: ")
    fmt.Println("Enter 5 fruits name: ")
    for i := 0;i<5;i++ {
      fmt.Printf("%v fruit:> ", i+1)
      fmt.Scan(&arr1[i]) 
    }
    fmt.Printf("You entered :")
    for in,j := range arr1 {
        fmt.Printf("%s",j)
        if(in!=len(arr1)-1) {
            fmt.Printf(", ")
        }
    }
    fmt.Println()
}
