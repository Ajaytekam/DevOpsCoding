package main

import (
    "fmt"
    "math"
)

func checkError(err error) {
    if err != nil {
        fmt.Println(err)
    }

}

func Sqrt(x float64) (float64, error) {

    if x < 0 {
        return 0, fmt.Errorf("cannot Sqrt negative number: %v", float64(x))
    }

    return math.Sqrt(x), nil
}

func main() {
    res1, err := Sqrt(2)
    checkError(err) 
    if res1 != 0 {fmt.Println(res1)}

    res2, err := Sqrt(-2)
    checkError(err) 
    if res2 != 0 {fmt.Println(res2)}
}
