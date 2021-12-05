package main

import (
    "fmt"
)

type Pair struct {
    First  int
    Second int
}

func Max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

// 博弈游戏：石头游戏
func stoneGame(piles []int) int {
    length := len(piles)
    if length == 0 {
        return 0
    } else if length == 1 {
        return piles[0]
    }
    dp := make([][]Pair, length, length)
    for i := 0; i < length; i++ {
        dp[i] = make([]Pair, length, length)
        dp[i][i].First = piles[i]
    }

    for i := length - 2; i >= 0; i-- {
        for j := i + 1; j < length; j++ {
            // [i,j]先手 = max(piles[i]+dp[i+1][j].Second,piles[j]+dp[i][j-1].Second)
            left := piles[i] + dp[i+1][j].Second
            right := piles[j] + dp[i][j-1].Second
            if left > right {
                dp[i][j].First = left
                dp[i][j].Second = dp[i+1][j].First
            } else {
                dp[i][j].First = right
                dp[i][j].Second = dp[i][j-1].First
            }
        }
    }
    last := &dp[0][length-1]
    return last.First - last.Second
}

func main() {
    fmt.Println(stoneGame([]int{1, 100, 3}))
    fmt.Println(stoneGame([]int{5, 3, 4, 5}))
}
