package main

func kthSmallest(root *TreeNode, k int) int {
    node := midOrder(root, &k)
    return node.Val
}

func midOrder(root *TreeNode, k *int) *TreeNode {
    if root == nil {
        return nil
    }

    node := midOrder(root.Left, k)
    if node != nil {
        return node
    }
    *k--
    if *k == 0 {
        return root
    }
    node = midOrder(root.Right, k)
    return node
}

func main() {

}
