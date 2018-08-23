package main

import (
    "fmt"
    "time"
)

func main() {
    counter := 0
    for {
        fmt.Println("MSB log", counter)
	counter += 1
        <-time.After(1 * time.Second)
    }
}
