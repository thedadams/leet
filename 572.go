/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    return isSubtreeRecurse(root, subRoot, true)
}

func isSubtreeRecurse(root *TreeNode, subRoot *TreeNode, recurse bool) bool {
    if root == nil {
        return subRoot == nil
    }
    if subRoot == nil {
        return false
    }

    if root.Val == subRoot.Val && isSubtreeRecurse(root.Left, subRoot.Left, false) && isSubtreeRecurse(root.Right, subRoot.Right, false) {
        return true
    }

    return recurse && (isSubtreeRecurse(root.Left, subRoot, true) || isSubtreeRecurse(root.Right, subRoot, true))
}