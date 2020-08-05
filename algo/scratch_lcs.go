package main

import (
    "fmt"
)

// 最长公共子序列
func lcs(text1 string, text2 string) int {
    l1, l2 := len(text1), len(text2)
    if l1 == 0 || l2 == 0 {
        return 0
    }
    dp := make([][]int, l1+1, l1+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, l2+1, l2+1)
    }
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                if dp[i-1][j] > dp[i][j-1] {
                    dp[i][j] = dp[i-1][j]
                } else {
                    dp[i][j] = dp[i][j-1]
                }
            }
        }
    }

    return dp[l1][l2]
}

func lcs2(text1 string, text2 string) int {
    l1, l2 := len(text1), len(text2)
    if l1 == 0 || l2 == 0 {
        return 0
    }

    if l1 < l2 {
        l1, l2 = l2, l1
        text1, text2 = text2, text1
    }

    dp := make([]int, l2+1, l2+1)
    for i := 1; i <= l1; i++ {
        prev := dp[0]
        for j := 1; j <= l2; j++ {
            if text1[i-1] == text2[j-1] {
                dp[j], prev = prev + 1, dp[j]
            } else {
                prev = dp[j]
                if dp[j] < dp[j-1] {
                    dp[j] = dp[j-1]
                }
            }
        }
    }

    return dp[l2]
}

func main() {
    fmt.Println(lcs("abcde", "ace"), lcs2("abcde", "ace"))
    fmt.Println(lcs("abc", "abc"), lcs2("abc", "abc"))
    fmt.Println(lcs("abc", "def"), lcs2("abc", "def"))
    fmt.Println(lcs("bsbininm", "jmjkbkjkv"), lcs2("bsbininm", "jmjkbkjkv"))
}
