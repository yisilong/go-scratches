package main

import (
    "fmt"
)

// 买卖股票问题
/*
   0<=n<=N-1  0<=k<=K
   dp[n][k].empty 表示第n天至多经过k次交易后，手里空仓的利润
   dp[n][k].full  表示第n天至多经过k次交易后，手里满仓的利润
   dp[n][k].empty = max(dp[n-1][k].empty, dp[n-1][k].full + prices[n])   // 昨天也空仓/昨天满仓&&今天卖出
   dp[n][k].full  = max(dp[n-1][k].full, dp[n-1][k-1].empty - prices[n]) // 昨天也满仓/昨天空仓&&今天买入
*/

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

type Equities struct {
    Empty int
    Full  int
}

// 最多1次交易
/*
   dp[n][1].empty = max(dp[n-1][1].empty, dp[n-1][1].full + prices[n])
   dp[n][1].full  = max(dp[n-1][1].full, dp[n-1][0].empty - prices[n])
                  = max(dp[n-1][1].full, -prices[n])
   化简后k不影响状态，再次简化
   dp[n].empty = max(dp[n-1].empty, dp[n-1].full + prices[n])
   dp[n].full  = max(dp[n-1].full, -prices[n])
*/

func maxProfit1(prices []int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    dp := make([]Equities, N, N)
    dp[0].Empty = 0
    dp[0].Full = -prices[0]
    for n := 1; n < N; n++ {
        dp[n].Empty = max(dp[n-1].Empty, dp[n-1].Full+prices[n])
        dp[n].Full = max(dp[n-1].Full, -prices[n])
    }
    return dp[N-1].Empty
}

func maxProfit1V2(prices []int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    pre := Equities{Empty: 0, Full: -prices[0]}
    for n := 1; n < N; n++ {
        tmp := pre
        pre.Empty = max(tmp.Empty, tmp.Full+prices[n])
        pre.Full = max(tmp.Full, -prices[n])
    }
    return pre.Empty
}

// 不限次数交易
/*
   数学上来看就是k=inf, k-1≈k, 可以省略k
   逻辑上来看就是没有这个状态限制, 当然就没有k什么事了
   dp[n].empty = max(dp[n-1].empty, dp[n-1].full + prices[n])
   dp[n].full  = max(dp[n-1].full, dp[n-1].empty - prices[n])
*/
func maxProfitInf(prices []int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    dp := make([]Equities, N, N)
    dp[0].Empty = 0
    dp[0].Full = -prices[0]
    for n := 1; n < N; n++ {
        dp[n].Empty = max(dp[n-1].Empty, dp[n-1].Full+prices[n])
        dp[n].Full = max(dp[n-1].Full, dp[n-1].Empty-prices[n])
    }
    return dp[N-1].Empty
}

func maxProfitInfV2(prices []int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    pre := Equities{Empty: 0, Full: -prices[0]}
    for n := 1; n < N; n++ {
        tmp := pre
        pre.Empty = max(tmp.Empty, tmp.Full+prices[n])
        pre.Full = max(tmp.Full, tmp.Empty-prices[n])
    }
    return pre.Empty
}

// 不限次数交易, 但卖出后需要冷却1天
/*
   第n天买入的时候, 要从n-2那天的状态转移, 不再是n-1
   dp[n].empty = max(dp[n-1].empty, dp[n-1].full + prices[n])
   dp[n].full  = max(dp[n-1].full, dp[n-2].empty - prices[n])
*/
func maxProfitInfWithCool(prices []int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    dp := make([]Equities, N, N)
    dp[0].Empty = 0
    dp[0].Full = -prices[0]
    prePreEmpty := 0
    for n := 1; n < N; n++ {
        dp[n].Empty = max(dp[n-1].Empty, dp[n-1].Full+prices[n])
        dp[n].Full = max(dp[n-1].Full, prePreEmpty-prices[n])
        prePreEmpty = dp[n-1].Empty
    }
    return dp[N-1].Empty
}

func maxProfitInfWithCoolV2(prices []int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    prePreEmpty := 0
    pre := Equities{Empty: 0, Full: -prices[0]}
    for n := 1; n < N; n++ {
        tmp := pre
        pre.Empty = max(tmp.Empty, tmp.Full+prices[n])
        pre.Full = max(tmp.Full, prePreEmpty-prices[n])
        prePreEmpty = tmp.Empty
    }
    return pre.Empty
}

// 不限次数交易, 但交易需要手续费
/*
   卖出或者买入把利润减掉，相当于买入的价格升高了
   dp[n].empty = max(dp[n-1].empty, dp[n-1].full + prices[n])
   dp[n].full  = max(dp[n-1].full, dp[n-1].empty - prices[n] - fee)
*/
func maxProfitInfWithFee(prices []int, fee int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    dp := make([]Equities, N, N)
    dp[0].Empty = 0
    dp[0].Full = -prices[0]
    for n := 1; n < N; n++ {
        dp[n].Empty = max(dp[n-1].Empty, dp[n-1].Full+prices[n])
        dp[n].Full = max(dp[n-1].Full, dp[n-1].Empty-prices[n]-fee)
    }
    return dp[N-1].Empty
}

func maxProfitInfWithFeeV2(prices []int, fee int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    pre := Equities{Empty: 0, Full: -prices[0]}
    for n := 1; n < N; n++ {
        tmp := pre
        pre.Empty = max(tmp.Empty, tmp.Full+prices[n])
        pre.Full = max(tmp.Full, tmp.Empty-prices[n]-fee)
    }
    return pre.Empty
}

// 最多k次交易
/*
   dp[n][k].empty = max(dp[n-1][k].empty, dp[n-1][k].full + prices[n])   // 昨天也空仓/昨天满仓&&今天卖出
   dp[n][k].full  = max(dp[n-1][k].full, dp[n-1][k-1].empty - prices[n]) // 昨天也满仓/昨天空仓&&今天买入
*/
func maxProfitK(prices []int, K int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    dp := make([][]Equities, N, N)
    for i := 0; i < N; i++ {
        dp[i] = make([]Equities, K+1, K+1)
        dp[i][0].Empty = 0
        dp[i][0].Full = 0
    }

    for k := 0; k <= K; k++ {
        dp[0][k].Empty = 0
        dp[0][k].Full = -prices[0]
    }

    for n := 1; n < N; n++ {
        for k := 1; k <= K; k++ {
            dp[n][k].Empty = max(dp[n-1][k].Empty, dp[n-1][k].Full+prices[n])
            dp[n][k].Full = max(dp[n-1][k].Full, dp[n-1][k-1].Empty-prices[n])
        }
    }
    return dp[N-1][K].Empty
}

func maxProfitKV2(prices []int, K int) int {
    N := len(prices)
    if N == 0 {
        return 0
    }
    if K > N/2 {
        return maxProfitInfV2(prices)
    }

    dp := make([][]Equities, N, N)
    for i := 0; i < N; i++ {
        dp[i] = make([]Equities, K+1, K+1)
        dp[i][0].Empty = 0
        dp[i][0].Full = 0
    }

    for k := 0; k <= K; k++ {
        dp[0][k].Empty = 0
        dp[0][k].Full = -prices[0]
    }

    for n := 1; n < N; n++ {
        for k := 1; k <= K; k++ {
            dp[n][k].Empty = max(dp[n-1][k].Empty, dp[n-1][k].Full+prices[n])
            dp[n][k].Full = max(dp[n-1][k].Full, dp[n-1][k-1].Empty-prices[n])
        }
    }
    return dp[N-1][K].Empty
}

func main() {
    prices := []int{2, 1, 4, 3, 6, 9, 10}
    fmt.Println(maxProfit1(prices), maxProfit1V2(prices))
    fmt.Println(maxProfitInf(prices), maxProfitInfV2(prices))
    fmt.Println(maxProfitInfWithCool(prices), maxProfitInfWithCoolV2(prices))
    fmt.Println(maxProfitInfWithFee(prices, 1), maxProfitInfWithFeeV2(prices, 1))
    fmt.Println(maxProfitK(prices, 1), maxProfitKV2(prices, 1))
    fmt.Println(maxProfitK(prices, 2), maxProfitKV2(prices, 2))
    fmt.Println(maxProfitK(prices, 4), maxProfitKV2(prices, 4))
}
