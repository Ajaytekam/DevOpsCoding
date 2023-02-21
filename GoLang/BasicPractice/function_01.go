package main 

import "fmt"  

func testFunc(x, y, z int) int {
    return x+y+z
}

func NewFunc1(x int, y int, msg string) (int, string) {
    mymessage := "Hello "+msg
    num := x + y
    return num, mymessage
}

func main() {
    fmt.Println(testFunc(10, 20, 30))
    ans, msg1 := NewFunc1(100, 100, "DonJoe")
    fmt.Println("Addition", ans) 
    fmt.Println("Message", msg1) 
}
