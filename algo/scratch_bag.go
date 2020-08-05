package main

import (
    "fmt"
)

// 01背包
func Bag(w int, wt []int, v []int) int {
    if w <= 0 || len(wt) <= 0 || len(wt) != len(v) {
        return 0
    }
    dp := make([][]int, len(wt)+1, len(wt)+1)
    for i := 0; i <= len(wt); i++ {
        dp[i] = make([]int, w+1, w+1)
    }

    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if j-wt[i-1] < 0 {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j-wt[i-1]] + v[i-1]
                if dp[i][j] < dp[i-1][j] {
                    dp[i][j] = dp[i-1][j]
                }
            }
        }
    }

    return dp[len(wt)][w]
}

func main() {
    fmt.Println(Bag(4, []int{2, 1, 2, 3}, []int{4, 2, 3, 6}))
}
