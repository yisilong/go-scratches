package main

import (
    "fmt"
)

type Node struct {
    Value string
    Next  *Node
}

// a->b->c
// a<-b<-c
func reverseList(head *Node) *Node {
    var pre, next *Node
    curr := head
    for curr != nil {
        next, curr.Next = curr.Next, pre
        pre, curr = curr, next
    }
    return pre
}

func makeList(arr []string) *Node {
    var head Node
    curr := &head
    for i := range arr {
        curr.Next = &Node{Value: arr[i]}
        curr = curr.Next
    }
    return head.Next
}

func Print(head *Node) {
    for head != nil {
        fmt.Printf("%s ", head.Value)
        head = head.Next
    }
    fmt.Println()
}

func main() {
    arr := []string{"a", "b", "c"}
    fmt.Println(arr)
    l := makeList(arr)
    Print(l)
    Print(reverseList(makeList([]string{})))
    Print(reverseList(makeList([]string{"a"})))
    Print(reverseList(makeList([]string{"a", "b"})))
    Print(reverseList(makeList([]string{"a", "b", "c"})))
    Print(reverseList(makeList([]string{"a", "b", "c", "d"})))
}
