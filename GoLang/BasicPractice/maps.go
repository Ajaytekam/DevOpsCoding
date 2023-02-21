package main  

import "fmt"

func main() {
    mymap := make(map[string]int)
    mymap["num1"] = 200 
    mymap["num2"] = 210 
    mymap["num3"] = 220 
    mymap["num4"] = 230 

    // print all maps 
    fmt.Println("Print all values : ")
    fmt.Println(mymap)

    // print key values  
    fmt.Println("\nPrint only keys:values pair : ")
    for key, val := range mymap {
        fmt.Println(key, "=>", val)
    }

    // print only keys  
    fmt.Println("\nPrint only keys : ")
    for key,_ := range mymap {
        fmt.Println(key)
    }
    
    // print only values   
    fmt.Println("\nPrint only values : ")
    for _,val := range mymap {
        fmt.Println(val)
    }
}
