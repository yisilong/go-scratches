package main

import (
    "gds/stack"
    "leetcode/helper"
)

type TreeNode = helper.TreeNode
type Stack = stack.ArrayStack

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
