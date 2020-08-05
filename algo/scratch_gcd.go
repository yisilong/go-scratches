package main

import (
    "fmt"
)

// 最大公约数
func gcd(n, m int) int {
    if m == 0 {
        return n
    } else {
        return gcd(m, n%m)
    }
}

func main() {
    fmt.Println(gcd(3, 6))
    fmt.Println(gcd(3, 7))
    fmt.Println(gcd(6, 3))
    fmt.Println(gcd(7, 3))
}
