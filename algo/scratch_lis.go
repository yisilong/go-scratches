package main

import (
    "fmt"
)

// 最长递增子序列
func longestLIS(arr []int) int {
    length := len(arr)
    if length <= 1 {
        return length
    }
    var ans int
    dp := make([]int, length, length)
    for i := 0; i < length; i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if arr[i] > arr[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        ans = max(ans, dp[i])
    }
    return ans
}

func longestLISV2(arr []int) int {
    length := len(arr)
    if length <= 1 {
        return length
    }
    var ans int
    top := make([]int, length, length)
    for i := 0; i < length; i++ {
        left, right := 0, ans-1
        for left <= right {
            mid := left + (right-left)/2
            if top[mid] >= arr[i] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
        top[left] = arr[i]
        if left == ans {
            ans++
        }
    }
    return ans
}

func main() {
    ans := []int{10, 9, 2, 5, 3, 7, 101, 18}
    fmt.Println(longestLIS(ans), longestLISV2(ans))
}
