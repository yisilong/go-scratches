package main

import (
    "fmt"
    "likebeta/leetcode/helper"
)

/*
   Morris 序
   1. 如果cur为null，则过程停止，否则继续下面的过程。
   2. 如果cur没有左子树，则让cur向右移动，即令cur=cur.right。
   3. 如果cur有左子树，则找到cur左子树上最右的节点，记为mostRight。
      a. 如果mostRight的right指针指向null，则令mostRight.right=cur，
         也就是让mostRight的right指针指向当前节点，然后让cur向左移动，即令cur=cur.left。
      b. 如果mostRight的right指针指向cur，则令mostRight.right=null，
         也就是让mostRight的right指针指向null，然后让cur向右移动，即令cur=cur.right。
*/
func Morris(head *TreeNode) {
    fmt.Println("Morris visited")
    curr := head
    for curr != nil {
        fmt.Printf("%v ", curr.Val)
        if curr.Left == nil {
            curr = curr.Right
            continue
        }
        leftMostRight := curr.Left
        for leftMostRight.Right != nil && leftMostRight.Right != curr {
            leftMostRight = leftMostRight.Right
        }
        if leftMostRight.Right == curr {
            leftMostRight.Right = nil
            curr = curr.Right
        } else {
            leftMostRight.Right = curr
            curr = curr.Left
        }
    }
    fmt.Println()
}

func PreOrderMorris(head *TreeNode) {
    fmt.Println("Morris PreOrder visited")
    curr := head
    for curr != nil {
        if curr.Left == nil {
            fmt.Printf("%v ", curr.Val)
            curr = curr.Right
            continue
        }
        leftMostRight := curr.Left
        for leftMostRight.Right != nil && leftMostRight.Right != curr {
            leftMostRight = leftMostRight.Right
        }
        if leftMostRight.Right == curr {
            leftMostRight.Right = nil
            curr = curr.Right
        } else {
            fmt.Printf("%v ", curr.Val)
            leftMostRight.Right = curr
            curr = curr.Left
        }
    }
    fmt.Println()
}

func InOrderMorris(head *TreeNode) {
    fmt.Println("Morris InOrder visited")
    curr := head
    for curr != nil {
        if curr.Left == nil {
            fmt.Printf("%v ", curr.Val)
            curr = curr.Right
            continue
        }
        leftMostRight := curr.Left
        for leftMostRight.Right != nil && leftMostRight.Right != curr {
            leftMostRight = leftMostRight.Right
        }
        if leftMostRight.Right == curr {
            fmt.Printf("%v ", curr.Val)
            leftMostRight.Right = nil
            curr = curr.Right
        } else {
            leftMostRight.Right = curr
            curr = curr.Left
        }
    }
    fmt.Println()
}

func PostOrderMorris(head *TreeNode) {
    fmt.Println("Morris PostOrder visited")
    // todo:
    fmt.Println()
}

func main() {
    t := helper.NewTree("[16,3,7,11,9,26,18,14,15]]")
    fmt.Println(t)
    Morris(t)
    PreOrderMorris(t)
    InOrderMorris(t)
}
