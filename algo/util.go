package main

import (
    "gds/stack"
    "leetcode/helper"
)

type ListNode = helper.ListNode
type TreeNode = helper.TreeNode
type Stack = stack.ArrayStack

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a int, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func IfElse(b bool, one, two int) int {
    if b {
        return one
    } else {
        return two
    }
}
