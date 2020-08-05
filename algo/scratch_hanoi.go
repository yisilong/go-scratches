package main

import (
    "fmt"
)

// 汉诺塔(不准直接左到右)
func hanoiProblem1(n int, side1, mid, side2 string) int {
    var step int
    if n >= 1 {
        step = process1(1, n, side1, mid, side2, side1, side2)
    }
    fmt.Println("total step ", step)
    return step
}

func process1(m, n int, side1, mid, side2 string, from, to string) int {
    if n == m {
        if from == mid || to == mid {
            fmt.Println(fmt.Sprintf("Move %d from %s to %s", n, from, to))
            return 1
        } else {
            fmt.Println(fmt.Sprintf("Move %d from %s to %s", n, from, mid))
            fmt.Println(fmt.Sprintf("Move %d from %s to %s", n, mid, to))
            return 2
        }
    }

    var step int
    if from == mid || to == mid {
        another := side1
        if from == side1 || to == side1 {
            another = side2
        }
        // 1. 将[m, n-1]从from移动到another
        step += process1(m, n-1, side1, mid, side2, from, another)
        // 2. 将n从from移动到to
        step += process1(n, n, side1, mid, side2, from, to)
        // 3. 将[m, n-1]从another移动到to
        step += process1(m, n-1, side1, mid, side2, another, to)
    } else {
        // 1. 将[m, n-1]从from移动到to
        step += process1(m, n-1, side1, mid, side2, from, to)
        // 2. 将n从from移动到mid
        step += process1(n, n, side1, mid, side2, from, mid)
        // 3. 将[m, n-1]从to移动到from
        step += process1(m, n-1, side1, mid, side2, to, from)
        // 4. 将n从mid移动到to
        step += process1(n, n, side1, mid, side2, mid, to)
        // 5. 将[m, n-1]从from移动到to
        step += process1(m, n-1, side1, mid, side2, from, to)
    }

    return step
}

func main() {
    hanoiProblem1(1, "left", "mid", "right")
    hanoiProblem1(2, "left", "mid", "right")
    hanoiProblem1(3, "left", "mid", "right")
}
