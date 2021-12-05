package main

import (
    "fmt"
    "leetcode/helper"
)

func maxBST(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    node, _1, _2, _3 := helper_(root)
    fmt.Println(_1, _2, _3)
    return node
}

func helper_(root *TreeNode) (*TreeNode, int, int, int) {
    if root.Left == nil && root.Right == nil {
        return root, 1, root.Val, root.Val
    }

    if root.Left == nil {
        rNode, rSize, rMin, rMax := helper_(root.Right)
        if rNode == root.Right && rMin > root.Val {
            return root, rSize + 1, root.Val, rMax
        } else {
            return rNode, rSize, rMin, rMax
        }
    } else if root.Right == nil {
        lNode, lSize, lMin, lMax := helper_(root.Left)
        if lNode == root.Left && lMax < root.Val {
            return root, lSize + 1, lMin, root.Val
        } else {
            return lNode, lSize, lMin, lMax
        }
    }
    lNode, lSize, lMin, lMax := helper_(root.Left)
    rNode, rSize, rMin, rMax := helper_(root.Right)
    if lNode == root.Left && rNode == root.Right && lMax < root.Val && rMin > root.Val {
        return root, lSize + rSize + 1, lMin, rMax
    }
    if lSize > rSize {
        return lNode, lSize, lMin, lMax
    } else {
        return rNode, rSize, rMin, rMax
    }
}

// 最大搜索二叉拓扑结构
type Share struct {
    Left  int
    Right int
}

type ShareMap map[*TreeNode]*Share

func maxBstTopo(root *TreeNode) int {
    cache := make(ShareMap)
    return postOrder(root, cache)
}

func max3(a, b, c int) int {
    if a < b {
        a = b
    }
    if a > c {
        return a
    } else {
        return c
    }
}

func postOrder(root *TreeNode, cache ShareMap) int {
    if root == nil {
        return 0
    }
    lMax := postOrder(root.Left, cache)
    rMax := postOrder(root.Right, cache)
    lCut := modifyCache(root.Left, root.Val, cache, true)
    rCut := modifyCache(root.Right, root.Val, cache, false)
    // var l, r int
    // if lShare, ok := cache[root.Left]; ok {
    //     l = lShare.Left + lShare.Right + 1
    // }
    // if rShare, ok := cache[root.Right]; ok {
    //     r = rShare.Left + rShare.Right + 1
    // }
    l, r := lMax-lCut, rMax-rCut
    cache[root] = &Share{Left: l, Right: r}
    return max3(l+r+1, lMax, rMax)
}

func modifyCache(root *TreeNode, v int, cache ShareMap, lr bool) int {
    if root == nil || cache[root] == nil {
        return 0
    }

    share := cache[root]
    if (lr && root.Val > v) || (!lr && root.Val < v) {
        delete(cache, root)
        return share.Left + share.Right + 1
    }

    var cutShare int
    if lr {
        cutShare = modifyCache(root.Right, v, cache, lr)
        share.Right -= cutShare
    } else {
        cutShare = modifyCache(root.Left, v, cache, lr)
        share.Left -= cutShare
    }
    return cutShare
}

func main() {
    t1 := helper.ParseTree("[6,1,12,0,3,10,13,null,null,null,null,4,14,20,16,2,5,11,15]")
    helper.PrintTree(t1, 8)
    node1 := maxBST(t1)
    helper.PrintTree(node1, 8)

    t2 := helper.ParseTree("[6,1,12,0,3,10,13,null,null,null,null,4,14,20,16,2,5,11,15]")
    helper.PrintTree(t2, 8)
    num := maxBstTopo(t2)
    fmt.Println(num)
}
