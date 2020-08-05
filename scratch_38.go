package main

import (
    "fmt"
)

func main() {
    count := 10
    for i := 0; i < count/2; i++ {
        fmt.Println(i)
        // count = 2
    }
}
