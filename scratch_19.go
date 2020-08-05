package main

import (
    "fmt"
)

func process(arr []int, i int , sum int, res map[int]bool) {
    if i == len(arr) {
        res[sum] = true
        return
    }

    process(arr, i+1, sum, res)
    process(arr, i+1, sum+arr[i], res)
}

func main() {
    res := make(map[int]bool)
    process([]int{1,2,3}, 0, 0, res)
    fmt.Println(res)
}
