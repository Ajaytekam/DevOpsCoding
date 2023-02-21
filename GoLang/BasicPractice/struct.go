package main

import "fmt"

type details struct {
    name string
    age int
    sex string
    address string
    occup string
}

func (dt details) Display() {
    fmt.Println("Name: ", dt.name)
    fmt.Println("Age: ", dt.age)
    fmt.Println("Sex: ", dt.sex)
    fmt.Println("Address: ", dt.address)
    fmt.Println("Occupation: ", dt.occup)
}
func main() {

    // declaring var method2  
    fmt.Println("Example 01: ")
    ob1 := new(details)

    ob1.name = "Ajay"
    ob1.age = 28
    ob1.sex = "male"
    ob1.address = "Ameerpet, Hydrabad"
    ob1.occup = "Job Seeker"

    // printing all the struct data 
    fmt.Println("Printing Values: ")
    fmt.Println(ob1.name)
    fmt.Println(ob1.age)
    fmt.Println(ob1.sex)
    fmt.Println(ob1.address)
    fmt.Println(ob1.occup)
    
    fmt.Println()
    
    // declaring var method3  
    fmt.Println("Example 02 using Structure Methods : ")
    var ob2 = &details{"Ajay", 28, "male", "Ameerpet Hydrabad", "Job Seeker"}
    fmt.Println("Printing Values: ")
    ob2.Display()
}
