package main

import (
    "fmt"
    "time"
    "sync"
)

func Print(wg *sync.WaitGroup, tm int, msg string) {
    time.Sleep(time.Duration(tm) * time.Millisecond)
    fmt.Println(msg)
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    go Print(&wg, 1000, "Hello world From 1") 
    go Print(&wg, 2000, "Hello world From 2") 
    go Print(&wg, 3000, "Hello world From 3") 
    wg.Wait()
    fmt.Println("All routines completed")
}
