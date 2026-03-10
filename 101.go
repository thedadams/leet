/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return root == nil || areSymmetric(root.Left, root.Right)
}

func areSymmetric(left, right *TreeNode) bool {
    if (left != nil) != (right != nil) {
        return false
    }

    if left == nil {
        return true
    }

    return left.Val == right.Val && areSymmetric(left.Left, right.Right) && areSymmetric(left.Right, right.Left)
}