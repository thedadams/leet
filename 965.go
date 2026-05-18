/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    return isUnival(root, root.Val)
}

func isUnival(root *TreeNode, val int) bool {
    if root.Val != val {
        return false
    }

    if root.Left != nil && !isUnival(root.Left, val) {
        return false
    }
    if root.Right != nil && !isUnival(root.Right, val) {
        return false
    }

    return true
}
