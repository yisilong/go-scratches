package main

import (
    "fmt"
)

type AVLNode struct {
    Left   *AVLNode
    Right  *AVLNode
    height int
    Value  interface{}
}

func (node *AVLNode) Height() int {
    if node == nil {
        return 0
    }
    return node.height
}

func (node *AVLNode) String() string {
    return fmt.Sprintf("%v/%d", node.Value, node.height)
}

func main() {
    n := AVLNode{height: 11, Value: "jljkljl"}
    fmt.Printf("%s\n", n.String())
}
