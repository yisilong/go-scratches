package main

import (
    "fmt"
)

type Node struct {
    Val  int
    Next *Node
}

func makeList(arr []int) *Node {
    var pt Node
    p := &pt
    for i := range arr {
        p.Next = &Node{Val: arr[i]}
        p = p.Next
    }
    return pt.Next
}

func makeArray(head *Node) []int {
    var arr []int
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    return arr
}

// 反转链表
func reverseList(head *Node) *Node {
    if head == nil || head.Next == nil {
        return head
    }
    last := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return last
}

// 反转链表前n个元素
func reverseListN(head *Node, n int) *Node {
    if head == nil || n <= 1 {
        return head
    }
    last, _ := funcReverseListN(head, n)
    return last
}

func funcReverseListN(head *Node, n int) (*Node, *Node) {
    if head.Next == nil || n <= 1 {
        return head, head.Next
    }
    last, lastNext := funcReverseListN(head.Next, n-1)
    head.Next.Next = head
    head.Next = lastNext
    return last, lastNext
}

// 反转链表第m个元素到第n个元素
func reverseListNM(head *Node, m int, n int) *Node {
    if head == nil || n-m < 1 || m < 1 {
        return head
    }
    return funcReverseListNM(head, m, n)
}

func funcReverseListNM(head *Node, m int, n int) *Node {
    if head.Next == nil || m == 1 {
        last, _ := funcReverseListN(head, n)
        return last
    }
    head.Next = funcReverseListNM(head.Next, m-1, n-1)
    return head
}

func main() {
    arr := []int{1, 3, 5, 9, 2}
    fmt.Println(makeArray(makeList(arr)))
    fmt.Println()
    //
    fmt.Println(makeArray(reverseList(makeList(arr))))
    fmt.Println()
    //
    fmt.Println(makeArray(reverseListN(makeList(arr), 0)))
    fmt.Println(makeArray(reverseListN(makeList(arr), 1)))
    fmt.Println(makeArray(reverseListN(makeList(arr), 2)))
    fmt.Println(makeArray(reverseListN(makeList(arr), 3)))
    fmt.Println(makeArray(reverseListN(makeList(arr), len(arr))))
    fmt.Println(makeArray(reverseListN(makeList(arr), len(arr)+1)))
    fmt.Println(makeArray(reverseListN(makeList(arr), len(arr)+2)))
    fmt.Println(makeArray(reverseListN(makeList(arr), len(arr)+10)))
    fmt.Println()
    //
    fmt.Println(makeArray(reverseListNM(makeList(arr), 0, 0)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 0, 2)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 1, 1)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 1, 2)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 1, 3)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 1, len(arr))))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 1, len(arr)+1)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 1, len(arr)+2)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 1, len(arr)+10)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 2, 3)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 2, 4)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 2, 5)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 2, len(arr))))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 2, len(arr)+1)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 2, len(arr)+2)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 2, len(arr)+3)))
    fmt.Println(makeArray(reverseListNM(makeList(arr), 2, len(arr)+10)))
}
