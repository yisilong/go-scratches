package main

import (
    "leetcode/helper"
)

// 打家劫舍

// 198.打家劫舍（简单）
/*
  不能偷相邻房间
*/
func rob1(nums []int) int {
    N := len(nums)
    dp := make([]int, N+2, N+2)
    for i := N - 1; i >= 0; i-- {
        dp[i] = max(dp[i+1], dp[i+2]+nums[i])
    }
    return dp[0]
}

func rob1V2(nums []int) int {
    N := len(nums)
    dpI, dpI1, dpI2 := 0, 0, 0
    for i := N - 1; i >= 0; i-- {
        dpI = max(dpI1, dpI2+nums[i])
        dpI2 = dpI1
        dpI1 = dpI
    }
    return dpI
}

// 213.打家劫舍II（中等）
/*
  不能偷相邻房间, 并且房子首尾相连
  首尾房间不能同时被抢，那么只可能有三种不同情况：要么都不被抢；要么第一间房子被抢最后一间不抢；要么最后一间房子被抢第一间不抢
  其实因为是求最大值，都不抢实际上包含着第二种和第三种情况中
*/
func rob2(nums []int) int {
    N := len(nums)
    if N == 0 {
        return 0
    } else if N == 1 {
        return nums[0]
    }

    return max(rob1V2(nums[:N-1]), rob1V2(nums[1:]))
}

// 337.打家劫舍III（中等）
/*
  打劫二叉树
*/
func rob3(root *TreeNode) int {
    rob, noRob := doRob(root)
    return max(rob, noRob)
}

func doRob(node *TreeNode) (int, int) {
    if node == nil {
        return 0, 0
    }
    robL, noRobL := doRob(node.Left)
    robR, noRobR := doRob(node.Right)
    return node.Val + noRobL + noRobR, max(robL, noRobL) + max(robR, noRobR)
}

func testRob(in string) {
    nums := helper.ParseArray(in)
    helper.Log(in, "=1=>", rob1(nums))
    helper.Log(in, "=1=>", rob1V2(nums))
    helper.Log(in, "=2=>", rob2(nums))
}

func testRob3(in string) {
    t := helper.ParseTree(in)
    helper.Log(in, "=3=>", rob3(t))
}

func main() {
    testRob("[1,2,3,1]")
    testRob("[2,7,9,3,1]")
    testRob3("[3,2,3,null,3,null,1]") // 7
    testRob3("[3,4,5,1,3,null,1]")    // 9
}
