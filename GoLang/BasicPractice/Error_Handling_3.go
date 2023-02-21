package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) (float64, error) {

    if x < 0 {
        return 0, fmt.Errorf("cannot Sqrt negative number: %v", float64(x))
    }

    return math.Sqrt(x), nil
}

func main() {
    res1, err := Sqrt(2)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(res1)
    }

    res2, err := Sqrt(-2)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(res2)
    }
}
