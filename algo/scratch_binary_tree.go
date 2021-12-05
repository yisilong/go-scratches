package main

import (
    "fmt"
    "leetcode/helper"
)

// 二叉树前序遍历
func preOrderVisit(head *TreeNode) {
    if head != nil {
        fmt.Printf("%d ", head.Val)
        preOrderVisit(head.Left)
        preOrderVisit(head.Right)
    }
}

func preOrderWithStack(head *TreeNode) {
    if head != nil {
        var s Stack
        s.Push(head)
        var curr *TreeNode
        for !s.Empty() {
            curr = s.Pop().(*TreeNode)
            fmt.Printf("%d ", curr.Val)
            if curr.Right != nil {
                s.Push(curr.Right)
            }
            if curr.Left != nil {
                s.Push(curr.Left)
            }
        }
    }
}

// 二叉树中序遍历
func inOrderVisit(head *TreeNode) {
    if head != nil {
        inOrderVisit(head.Left)
        fmt.Printf("%d ", head.Val)
        inOrderVisit(head.Right)
    }
}

func inOrderWithStack1(head *TreeNode) {
    if head != nil {
        var s Stack
        curr := head
        for curr != nil || !s.Empty() {
            if curr != nil {
                s.Push(curr)
                curr = curr.Left
            } else {
                curr = s.Pop().(*TreeNode)
                fmt.Printf("%d ", curr.Val)
                curr = curr.Right
            }
        }
    }
}

// 二叉树后序遍历
func postOrderVisit(head *TreeNode) {
    if head != nil {
        postOrderVisit(head.Left)
        postOrderVisit(head.Right)
        fmt.Printf("%d ", head.Val)
    }
}

func postOrderWithStack1(head *TreeNode) {
    if head != nil {
        var s Stack
        s.Push(head)
        var curr, pre *TreeNode
        for !s.Empty() {
            curr = s.Peek().(*TreeNode)
            if curr.Left != nil && pre != curr.Left && pre != curr.Right {
                s.Push(curr.Left)
            } else if curr.Right != nil && pre != curr.Right {
                s.Push(curr.Right)
            } else {
                fmt.Printf("%d ", curr.Val)
                pre = s.Pop().(*TreeNode)
            }
        }
    }
}

func postOrderWithStack2(head *TreeNode) {
    if head != nil {
        var s1, s2 Stack
        s1.Push(head)
        for !s1.Empty() {
            curr := s1.Pop().(*TreeNode)
            s2.Push(curr)
            if curr.Left != nil {
                s1.Push(curr.Left)
            }
            if curr.Right != nil {
                s1.Push(curr.Right)
            }
        }
        for !s2.Empty() {
            fmt.Printf("%d ", s2.Pop().(*TreeNode).Val)
        }
    }
}

func main() {
    arr := "[5,2,1,4,3,6,7]"
    fmt.Println(arr)
    t := helper.ParseTree(arr)
    fmt.Println(t)

    fmt.Println("前序")
    preOrderVisit(t)
    fmt.Println()
    preOrderWithStack(t)
    fmt.Println()

    fmt.Println("中序")
    inOrderVisit(t)
    fmt.Println()
    inOrderWithStack1(t)
    fmt.Println()

    fmt.Println("后序")
    postOrderVisit(t)
    fmt.Println()
    postOrderWithStack1(t)
    fmt.Println()
    postOrderWithStack2(t)
    fmt.Println()
}
