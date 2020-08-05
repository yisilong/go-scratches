package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node, _ := recurseTree(root, p, q)
	return node
}

func recurseTree(node, p, q *TreeNode) (*TreeNode, bool) {
	if node == nil {
		return nil, false
	}

	var count int

	left, leftExist := recurseTree(node.Left, p, q)
	if left != nil {
		return left, true
	}

	if leftExist {
		count++
	}

	var midExist bool
	if node == p || node == q {
		midExist = true
	}

	if midExist {
		count++
	}

	if count == 2 {
		return node, true
	}

	right, rightExist := recurseTree(node.Right, p, q)
	if right != nil {
		return right, true
	}

	if rightExist {
		count++
	}

	if count == 2 {
		return node, true
	}

	return nil, count > 0
}

func main() {

}
