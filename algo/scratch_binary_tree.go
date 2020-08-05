package main

import (
    "fmt"
    "local/gds"
)

type Node = gds.TreeNode
type Stack = gds.Stack

// 二叉树前序遍历
func preOrderVisit(head *Node) {
    if head != nil {
        fmt.Printf("%d ", head.Value)
        preOrderVisit(head.Left)
        preOrderVisit(head.Right)
    }
}

func preOrderWithStack(head *Node) {
    if head != nil {
        var s Stack
        s.Push(head)
        var curr *Node
        for !s.Empty() {
            curr = s.Pop().(*Node)
            fmt.Printf("%d ", curr.Value)
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
func inOrderVisit(head *Node) {
    if head != nil {
        inOrderVisit(head.Left)
        fmt.Printf("%d ", head.Value)
        inOrderVisit(head.Right)
    }
}

func inOrderWithStack1(head *Node) {
    if head != nil {
        var s Stack
        curr := head
        for curr != nil || !s.Empty() {
            if curr != nil {
                s.Push(curr)
                curr = curr.Left
            } else {
                curr = s.Pop().(*Node)
                fmt.Printf("%d ", curr.Value)
                curr = curr.Right
            }
        }
    }
}

// 二叉树后序遍历
func postOrderVisit(head *Node) {
    if head != nil {
        postOrderVisit(head.Left)
        postOrderVisit(head.Right)
        fmt.Printf("%d ", head.Value)
    }
}

func postOrderWithStack1(head *Node) {
    if head != nil {
        var s Stack
        s.Push(head)
        var curr, pre *Node
        for !s.Empty() {
            curr = s.Peek().(*Node)
            if curr.Left != nil && pre != curr.Left && pre != curr.Right {
                s.Push(curr.Left)
            } else if curr.Right != nil && pre != curr.Right {
                s.Push(curr.Right)
            } else {
                fmt.Printf("%d ", curr.Value)
                pre = s.Pop().(*Node)
            }
        }
    }
}

func postOrderWithStack2(head *Node) {
    if head != nil {
        var s1, s2 Stack
        s1.Push(head)
        for !s1.Empty() {
            curr := s1.Pop().(*Node)
            s2.Push(curr)
            if curr.Left != nil {
                s1.Push(curr.Left)
            }
            if curr.Right != nil {
                s1.Push(curr.Right)
            }
        }
        for !s2.Empty() {
            fmt.Printf("%d ", s2.Pop().(*Node).Value)
        }
    }
}

func main() {
    arr := []int{5, 2, 1, 4, 3, 6, 7}
    // rand.Shuffle(len(arr), func(i, j int) {
    //     arr[i], arr[j] = arr[j], arr[i]
    // })
    fmt.Println(arr)
    t := gds.NewTree(arr)
    fmt.Println(t)

    fmt.Println("前序")
    preOrderVisit(t.Root)
    fmt.Println()
    preOrderWithStack(t.Root)
    fmt.Println()

    fmt.Println("中序")
    inOrderVisit(t.Root)
    fmt.Println()
    inOrderWithStack1(t.Root)
    fmt.Println()

    fmt.Println("后序")
    postOrderVisit(t.Root)
    fmt.Println()
    postOrderWithStack1(t.Root)
    fmt.Println()
    postOrderWithStack2(t.Root)
    fmt.Println()
}
