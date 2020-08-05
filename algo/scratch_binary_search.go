package main

import (
    "fmt"
)

// 二分搜索
func binarySearch(arr []int, target int) int {
    length := len(arr)
    if length == 0 {
        return -1
    }
    left := 0
    right := length - 1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else { // arr[mid] > target
            right = mid - 1
        }
    }
    return -1
}

// 二分搜索最左出现
func leftBound(arr []int, target int) int {
    length := len(arr)
    if length == 0 {
        return -1
    }
    left := 0
    right := length - 1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == target {
            right = mid - 1
        } else if arr[mid] < target {
            left = mid + 1
        } else { // arr[mid] > target
            right = mid - 1
        }
    }

    // 搜索区间[left, right] =>[right, left] 即 [left-1, left]
    if left >= length || arr[left] != target {
        return -1
    }
    return left
}

// 二分搜索最右出现
func rightBound(arr []int, target int) int {
    length := len(arr)
    if length == 0 {
        return -1
    }
    left := 0
    right := length - 1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == target {
            left = mid + 1
        } else if arr[mid] < target {
            left = mid + 1
        } else { // arr[mid] > target
            right = mid - 1
        }
    }

    // 搜索区间[left, right] =>[right, left] 即 [right, right+1]
    if right <= -1 || arr[right] != target {
        return -1
    }
    return right
}

func main() {
    arr := []int{1, 2, 2, 2, 3, 3, 4, 5}
    for i := 0; i < 10; i++ {
        a := binarySearch(arr, i)
        b := leftBound(arr, i)
        c := rightBound(arr, i)
        fmt.Printf("%d: %2d %2d %2d\n", i, a, b, c)
    }
    fmt.Println([]int{1,2}[:0])
}
