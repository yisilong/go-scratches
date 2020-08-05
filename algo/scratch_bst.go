package main

import (
    "fmt"
	"math/rand"
	"strings"
)

// --------------------------------------------------------------------------------------------------------------------
// 二叉搜索树
type BSTNode struct {
	Left  *BSTNode
	Right *BSTNode
	Value int
}

type BSTree = BSTNode

func isValid(root *BSTNode, min *BSTNode, max *BSTNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Value <= min.Value {
		return false
	} else if max != nil && root.Value >= max.Value {
		return false
	}
	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}

func IsValid(root *BSTree) bool {
	return isValid(root, nil, nil)
}

func Add(root *BSTree, value int) *BSTree {
	if root == nil {
		return &BSTNode{Value: value}
	}
	if root.Value > value {
		root.Left = Add(root.Left, value)
	} else if root.Value < value {
		root.Right = Add(root.Right, value)
	}
	return root
}

func Find(root *BSTree, value int) *BSTNode {
	if root == nil {
		return nil
	}

	if root.Value == value {
		return root
	} else if root.Value > value {
		return Find(root.Left, value)
	} else {
		return Find(root.Right, value)
	}
}

func getLeftMax(node *BSTNode) (*BSTNode, *BSTNode) {
	parent, curr := node, node.Left
	for curr.Right != nil {
		parent = curr
		curr = curr.Right
	}
	return parent, curr
}

func Delete(root *BSTree, value int) *BSTree {
	if root == nil {
		return nil
	}
	if root.Value > value {
		root.Left = Delete(root.Left, value)
	} else if root.Value < value {
		root.Right = Delete(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			parent, node := getLeftMax(root)
			root.Value = node.Value
			// root.Left = Delete(root.Left, node.Value)
			if parent == root {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		}
	}
	return root
}

// --------------------------------------------------------------------------------------------------------------------

func Print(root *BSTree) {
	fmt.Println("Binary Search Tree:")
	printInOrder(root, 0, "H", 17)
	fmt.Println()
}

func printInOrder(head *BSTNode, height int, to string, length int) {
	if head == nil {
		return
	}
	printInOrder(head.Right, height+1, "v", length)
	val := fmt.Sprintf("%s%d%s", to, head.Value, to)
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val = strings.Repeat(" ", lenL) + val + strings.Repeat(" ", lenR)
	fmt.Println(strings.Repeat(" ", height*length) + val)
	printInOrder(head.Left, height+1, "^", length)
}

func newTree(arr []int) *BSTree {
	var head *BSTNode
	for i := 0; i < len(arr); i++ {
		head = Add(head, arr[i])
	}
	return head
}

func main() {
	arr := []int{5, 2, 1, 4, 3, 6, 7}
	t := newTree(arr)
	fmt.Println(IsValid(t))
	Print(t)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)
	for i := range arr {
		fmt.Println("delete ", arr[i])
		t = Delete(t, arr[i])
		Print(t)
	}
}
