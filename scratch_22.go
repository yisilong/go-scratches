package main

import (
    "fmt"
)

func diff(a, b int) bool {
    return a^b < 0
}

func main() {
    fmt.Println(int('a'), int('A'))
    fmt.Println('a'|' ', 'A'|' ')
    fmt.Println('a'&'_', 'A'&'_')
    fmt.Println('a'^' ', 'A'^' ')
    fmt.Println(diff(5, -8), diff(-6, 3), diff(-2, -5), diff(8, 7))
    n := 3
    fmt.Println(n, -^n, ^-n)
}
