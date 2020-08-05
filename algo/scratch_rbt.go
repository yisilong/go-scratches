package main

import (
    "fmt"
	"math/rand"
	"strings"
)

// --------------------------------------------------------------------------------------------------------------------
// 红黑树
/*
	一棵二叉查找树为一棵红黑树，需要满足下面的性质：
	1. 每个结点是红的，或是黑的
	2. 根结点是黑的
	3. 每个叶结点（NIL）是黑的
	4. 如果一个结点是红的，则它的两个儿子都是黑的
	5. 对每个结点，从该结点到其子孙结点的所有路径上包含相同数目的黑结点
*/
const (
	RED   = true
	BLACK = false
)

type RBTNode struct {
	Left   *RBTNode
	Right  *RBTNode
	Parent *RBTNode
	Color  bool
	Value  int
}

func (node *RBTNode) IsRed() bool {
	return node != nil && node.Color == RED
}

func (node *RBTNode) IsBlack() bool {
	return node == nil || node.Color == BLACK
}

type RBTree struct {
	root *RBTNode
	size int
	// todo: value comparator
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (t *RBTree) Add(value int) *RBTree {
	var insertedNode *RBTNode
	if t.root == nil {
		t.root = &RBTNode{Value: value, Color: RED}
		insertedNode = t.root
	} else {
		for node := t.root; ; {
			if node.Value == value {
				return t
			} else if node.Value > value {
				if node.Left != nil {
					node = node.Left
				} else {
					node.Left = &RBTNode{Parent: node, Value: value, Color: RED}
					insertedNode = node.Left
					break
				}
			} else {
				if node.Right != nil {
					node = node.Right
				} else {
					node.Right = &RBTNode{Parent: node, Value: value, Color: RED}
					insertedNode = node.Right
					break
				}
			}
		}
	}
	t.size++
	t.insertCase1(insertedNode)
	return t
}

/*
	N为红黑树根结点, 直接标记为黑色
		N --->  Nb
*/
func (t *RBTree) insertCase1(node *RBTNode) {
	if node.Parent == nil {
		node.Color = BLACK
	} else {
		t.insertCase2(node)
	}
}

/*
	N的父节点P是黑色, 不需要调整
		Pb               Pb
       /                   \
      Nr                    Nr
*/
func (t *RBTree) insertCase2(node *RBTNode) {
	if node.Parent.IsBlack() {
		// do nothing
	} else {
		t.insertCase3(node)
	}
}

/*
	N的父节点P是红色并且叔叔U也是红色, 向上递归调整颜色, N可以是P的左右节点, P可以是G的左右节点，情况一样
		  Gb             Gb             Gr
         /  \           /  \           /  \
		Pr  Ur         Pb  Ub         Pb  Ub
       /              /              /
      Nr             Nr             Nr
*/
func (t *RBTree) insertCase3(node *RBTNode) {
	uncle := node.Parent.Parent.Left
	if node.Parent == node.Parent.Parent.Left {
		uncle = node.Parent.Parent.Right
	}
	if uncle.IsRed() {
		node.Parent.Color = BLACK
		uncle.Color = BLACK
		uncle.Parent.Color = RED
		t.insertCase1(uncle.Parent)
	} else {
		t.insertCase4(node)
	}
}

/*
	N的父节点P是红色并且叔叔U是黑色,  对PN左旋或者右旋变形成case5
		  Gb             Gb            |            Gb             Gb
         /  \           /  \           |           /  \           /  \
		Pr  Ub         Nr  Ub          |          Ub  Pr         Ub  Nr
         \            /  \             |              /             /  \
         Nr          Pr   B            |             Nr            A   Pr
		/  \          \                |            /  \               /
       A    B          A               |           A    B             B
*/
func (t *RBTree) insertCase4(node *RBTNode) {
	grandfather := node.Parent.Parent
	if node == node.Parent.Right && node.Parent == grandfather.Left {
		grandfather.Left = t.leftRotate(node.Parent)
		node = grandfather.Left.Left
	} else if node == node.Parent.Left && node.Parent == grandfather.Right {
		grandfather.Right = t.rightRotate(node.Parent)
		node = grandfather.Right.Right
	}
	t.insertCase5(node)
}

/*
	N的父节点P是红色并且叔叔U是黑色, 对GP右旋或者左旋
		  Gb              Pb            |          Gb             Pb
         /  \            /  \           |         /  \           /  \
		Pr  Ub          Nr  Gr          |        Ub  Pr         Gr  Nr
       /  \                /  \         |           /  \       /  \
      Nr   C              C   Ub        |          C   Nr     Ub   C
*/
func (t *RBTree) insertCase5(node *RBTNode) {
	grandfather := node.Parent.Parent
	node.Parent.Color = BLACK
	grandfather.Color = RED
	if node == node.Parent.Left && node.Parent == grandfather.Left {
		t.checkRoot(grandfather.Left, grandfather)
		t.rightRotate(grandfather)
	} else if node == node.Parent.Right && node.Parent == grandfather.Right {
		t.checkRoot(grandfather.Right, grandfather)
		t.leftRotate(grandfather)
	}
}

func (t *RBTree) checkRoot(newRoot *RBTNode, oldRoot *RBTNode) {
	if oldRoot == t.root {
		t.root = newRoot
	} else if oldRoot.Parent.Left == oldRoot {
		oldRoot.Parent.Left = newRoot
	} else /*if oldRoot.Parent.Right == oldRoot*/ {
		oldRoot.Parent.Right = newRoot
	}
}

func (t *RBTree) leftRotate(node *RBTNode) *RBTNode {
	return rrAdjust(node)
}

func (t *RBTree) rightRotate(node *RBTNode) *RBTNode {
	return llAdjust(node)
}

/*
  		  A                  B
         /                  / \
        B                  C   A
       / \                    /
      C   D                  D
*/
func llAdjust(a *RBTNode) *RBTNode {
	b := a.Left
	a.Left = b.Right // d = b.Right
	if b.Right != nil {
		b.Right.Parent = a
	}
	b.Right = a
	b.Parent = a.Parent
	a.Parent = b
	return b
}

/*
  	  A                      B
       \                    / \
        B                  A   D
       / \                  \
      C   D                  C
*/
func rrAdjust(a *RBTNode) *RBTNode {
	b := a.Right
	a.Right = b.Left // c = b.Left
	if b.Left != nil {
		b.Left.Parent = a
	}
	b.Left = a
	b.Parent = a.Parent
	a.Parent = b
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
func lrAdjust(a *RBTNode) *RBTNode {
	b := a.Left
	c := b.Right
	a.Left = c.Right // e = c.Right
	if c.Right != nil {
		c.Right.Parent = a
	}
	b.Right = c.Left // d = c.Left
	if c.Left != nil {
		c.Left.Parent = b
	}
	c.Left = b
	c.Right = a
	c.Parent = a.Parent
	a.Parent = c
	b.Parent = c
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
func rlAdjust(a *RBTNode) *RBTNode {
	b := a.Right
	c := b.Left
	a.Right = c.Left // d = c.Left
	if c.Left != nil {
		c.Left.Parent = a
	}
	b.Left = c.Right // e = c.Right
	if c.Right != nil {
		c.Right.Parent = b
	}
	c.Left = a
	c.Right = b
	c.Parent = a.Parent
	a.Parent = c
	b.Parent = c
	return c
}

func (t *RBTree) Find(value int) bool {
	node := t.root
	for node != nil {
		if node.Value == value {
			break
		} else if node.Value > value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node == nil
}

func (t *RBTree) Delete(value int) *RBTree {
	// todo:
	return t
}

func (t *RBTree) Print() {
	fmt.Println("Red Block Tree:")
	t.printInOrder(t.root, 0, "H", 17)
	fmt.Println()
}

func (t *RBTree) printInOrder(node *RBTNode, height int, to string, length int) {
	if node != nil {
		t.printInOrder(node.Right, height+1, "v", length)
		var val string
		if node.IsRed() {
			val = fmt.Sprintf("%s%d/b%s", to, node.Value, to)
		} else {
			val = fmt.Sprintf("%s%d/d%s", to, node.Value, to)
		}
		lenM := len(val)
		lenL := (length - lenM) / 2
		lenR := length - lenM - lenL
		val = strings.Repeat(" ", lenL) + val + strings.Repeat(" ", lenR)
		fmt.Println(strings.Repeat(" ", height*length) + val)
		t.printInOrder(node.Left, height+1, "^", length)
	}
}

// --------------------------------------------------------------------------------------------------------------------
func newTree(arr []int) *RBTree {
	rbt := NewRBTree()
	for i := 0; i < len(arr); i++ {
		rbt.Add(arr[i])
	}
	return rbt
}

func main() {
	arr := []int{13, 8, 17, 1, 11, 15, 25, 6, 22, 27}
	t := newTree(arr)
	t.Print()
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)
	for i := range arr {
		fmt.Println("delete ", arr[i])
		t.Print()
	}
}
