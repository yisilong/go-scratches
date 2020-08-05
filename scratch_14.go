package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var list []int
    list = make([]int, 12, 12)
    fmt.Println(list)
    s := "[1,24,5]"
    err := json.Unmarshal([]byte(s), &list)
    fmt.Println(list, err)
}
