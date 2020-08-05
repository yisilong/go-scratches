package main

import (
    "fmt"
	"math/rand"
	"strings"
)

// --------------------------------------------------------------------------------------------------------------------
// 平衡二叉树
type AVLNode struct {
	Left   *AVLNode
	Right  *AVLNode
	Height int
	Value  int
}

type AVLTree = AVLNode

func Add(root *AVLTree, value int) *AVLTree {
	if root == nil {
		return &AVLNode{Value: value, Height: 1}
	}
	if root.Value > value {
		root.Left = Add(root.Left, value)
		root = adjust(root)
	} else if root.Value < value {
		root.Right = Add(root.Right, value)
		root = adjust(root)
	}
	return root
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func updateHeight(node *AVLNode) {
	node.Height = max(getHeight(node.Left), getHeight(node.Right)) + 1
}

func adjust(node *AVLNode) *AVLNode {
	lH, rH := getHeight(node.Left), getHeight(node.Right)
	if lH-rH == 2 {
		if getHeight(node.Left.Left)-getHeight(node.Left.Right) > 0 {
			return llAdjust(node)
		} else {
			return lrAdjust(node)
		}
	} else if lH-rH == -2 {
		if getHeight(node.Right.Left)-getHeight(node.Right.Right) < 0 {
			return rrAdjust(node)
		} else {
			return rlAdjust(node)
		}
	} else {
		node.Height = max(lH, rH) + 1
	}
	return node
}

/*
  		  A                  B
         /                  / \
        B                  C   A
       / \                    /
      C   D                  D
*/
func llAdjust(node *AVLNode) *AVLNode {
	b := node.Left
	node.Left = b.Right
	b.Right = node
	updateHeight(node)
	updateHeight(b)
	return b
}

/*
  	  A                      B
       \                    / \
        B                  A   D
       / \                  \
      C   D                  C
*/
func rrAdjust(node *AVLNode) *AVLNode {
	b := node.Right
	node.Right = b.Left
	b.Left = node
	updateHeight(node)
	updateHeight(b)
	return b
}

/*
		A
       /                   C
      B                  /   \
       \                B     A
        C                \   /
       / \                D E
      D   E
*/
func lrAdjust(node *AVLNode) *AVLNode {
	b := node.Left
	c := b.Right
	node.Left = c.Right
	b.Right = c.Left
	c.Left = b
	c.Right = node
	updateHeight(node)
	updateHeight(b)
	updateHeight(c)
	return c
}

/*
     A
      \                   C
       B                /   \
      /                A     B
     C                  \   /
    / \                  D E
   D   E
*/
func rlAdjust(node *AVLNode) *AVLNode {
	b := node.Right
	c := b.Left
	node.Right = c.Left
	b.Left = c.Right
	c.Left = node
	c.Right = b
	updateHeight(node)
	updateHeight(b)
	updateHeight(c)
	return c
}

func Find(root *AVLTree, value int) *AVLNode {
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

func getLeftMax(node *AVLNode) *AVLNode {
	curr := node.Left
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr
}

func getRightMin(node *AVLNode) *AVLNode {
	curr := node.Right
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func Delete(root *AVLTree, value int) *AVLTree {
	if root == nil {
		return nil
	}
	if root.Value > value {
		root.Left = Delete(root.Left, value)
	} else if root.Value < value {
		root.Right = Delete(root.Right, value)
	} else {
		if root.Left == nil {
			root = root.Right
		} else if root.Right == nil {
			root = root.Left
		} else {
			if getHeight(root.Left) > getHeight(root.Right) {
				node := getLeftMax(root)
				root.Value = node.Value
				root.Left = Delete(root.Left, node.Value)
			} else {
				node := getRightMin(root)
				root.Value = node.Value
				root.Right = Delete(root.Right, node.Value)
			}
		}
	}
	if root != nil {
		root = adjust(root)
	}
	return root
}

// --------------------------------------------------------------------------------------------------------------------

func Print(root *AVLTree) {
	fmt.Println("AVL Tree:")
	printInOrder(root, 0, "H", 17)
	fmt.Println()
}

func printInOrder(head *AVLNode, height int, to string, length int) {
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

func newTree(arr []int) *AVLTree {
	var head *AVLNode
	for i := 0; i < len(arr); i++ {
		head = Add(head, arr[i])
	}
	return head
}

func main() {
	arr := []int{16, 3, 7, 11, 9, 26, 18, 14, 15}
	t := newTree(arr)
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
