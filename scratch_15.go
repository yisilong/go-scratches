package main

import (
    "fmt"
    "github.com/likebeta/gds/tree"
    "github.com/likebeta/gds/util"
    "math/rand"
)

func main() {
    arr := []int{5, 2, 1, 4, 3, 6, 7}
    rand.Shuffle(len(arr), func(i, j int) {
        arr[i], arr[j] = arr[j], arr[i]
    })
    fmt.Println(arr)
    t := tree.NewBSTree(util.IntComparator)
    for i := range arr {
        t.Add(arr[i])
    }
    fmt.Println(t.IsValid(), t)
    for i := range arr {
        fmt.Println("delete ", arr[i])
        t = t.Delete(arr[i])
        fmt.Println(t.IsValid(), t)
    }
}
