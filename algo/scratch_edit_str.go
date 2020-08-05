package main

import (
    "fmt"
)

// 最少编辑距离
func minDistance(word1 string, word2 string) int {
    lS, lD := len(word1), len(word2)
    if lS == 0 {
        return lD
    } else if lD == 0 {
        return lS
    }

    dp := make([][]int, lS+1, lS+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, lD+1, lD+1)
        dp[i][0] = i
    }

    for j := 0; j < len(dp[0]); j++ {
        dp[0][j] = j
    }

    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            dp[i][j] = dp[i-1][j-1]
            if word1[i-1] != word2[j-1] {
                if dp[i][j] > dp[i][j-1] {
                    dp[i][j] = dp[i][j-1]
                }
                if dp[i][j] > dp[i-1][j] {
                    dp[i][j] = dp[i-1][j]
                }
                dp[i][j] += 1
            }
        }
    }

    return dp[lS][lD]
}

func main() {
    fmt.Println(minDistance("", "23"))
    fmt.Println(minDistance("11", ""))
    fmt.Println(minDistance("1126", "216"))
    fmt.Println(minDistance("horse", "ros"))
    fmt.Println(minDistance("intention", "execution"))
}
