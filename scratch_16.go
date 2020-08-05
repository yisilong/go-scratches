package main

import (
    "fmt"
)

func main() {
    s := make(map[int]string)
    s[1] = "hi"
    s[2] = "what"
    fmt.Println(s)
    delete(s, 5)
    fmt.Println(s)
}
