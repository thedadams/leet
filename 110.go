/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, ok := isHeightBalanced(root)
    return ok
}

func isHeightBalanced(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }

    leftHeight, ok := isHeightBalanced(root.Left)
    if !ok {
        return 0, ok
    }

    rightHeight, ok := isHeightBalanced(root.Right)
    if !ok {
        return 0, ok
    }

    height := max(leftHeight, rightHeight)
    return height + 1, height - min(leftHeight, rightHeight) <= 1
}