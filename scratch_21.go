package main

import (
    "fmt"
)

type Node struct {
    Value int
    Next  *Node
}

func reverseList(node *Node) *Node {
    if node == nil || node.Next == nil {
        return node
    }

    newList := reverseList(node.Next)
    node.Next.Next = node
    node.Next = nil
    return newList
}

func newList(arr []int) *Node {
    head := Node{}
    p := &head
    for i := range arr {
        p.Next = &Node{Value: arr[i]}
        p = p.Next
    }
    return head.Next
}

func toArr(head *Node) []int {
    var arr []int
    for head != nil {
        arr = append(arr, head.Value)
        head = head.Next
    }
    return arr
}

func testOne(arr []int) {
    list := newList(arr)
    listArr := toArr(list)
    rList := reverseList(list)
    rListArr := toArr(rList)
    fmt.Printf("arr: %v, listArr: %v, rListArr: %v\n", arr, listArr, rListArr)
}

func main() {
    testOne(nil)
    testOne([]int{})
    testOne([]int{1})
    testOne([]int{1,2})
    testOne([]int{1,2, 6})
    testOne([]int{1,2, 6, 8, 9})
    var i int32 = 1<<31 - 1
    var j int32 = int32(^uint32(0) >> 1)
    // var j int32 = 1<<32
    fmt.Println(i, j, i+(j-i)/2)
}
