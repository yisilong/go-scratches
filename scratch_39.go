package main

import (
    "fmt"
)

func ttt() (int, int) {
    return 2, 3
}

func Ppp() (ans string) {
    ans = "hello"
    return ans + "world"
}

func main() {
    var i, num int
    if i > -10 {
        i, num = ttt()
        fmt.Println(i, num)
    }
    fmt.Println(i, num)
    fmt.Println("hello"[5:5])
    fmt.Println(Ppp())
}
