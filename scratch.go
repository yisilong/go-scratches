package main

import (
    "fmt"
)

func add(arr []int) []int {
    arr = append(arr, 10)
    fmt.Printf("%v %p\n", arr, &arr)
    return arr
}

func add1(arr []int, result [][]int) [][]int {
    arr = append(arr, 5)
    return append(result, arr)
}

func add2(arr []int, result [][]int) [][]int {
    arr = append(arr, 9)
    return append(result, arr)
}

func main() {
    // var rcs []int
    // fmt.Println(rcs)
    // rcs = append(rcs, 111)
    // fmt.Println(rcs)
    // rcs = nil
    // fmt.Println(rcs)
    // rcs = append(rcs, 111)
    // fmt.Println(rcs)

    var result [][]int
    arr := []int{1}
    fmt.Printf("%v %p\n", arr, &arr)
    result = add1(arr, result)
    arr3 := arr[:]
    arr3 = append(arr3, 89)
    arr[0] = 100
    result = add2(arr, result)
    fmt.Printf("%v %p\n", arr, &arr)
    result = append(result, arr3)
    fmt.Printf("%v\n", result)
}
