/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    return findSumTarget(root, k, make(map[int]struct{}))
}

func findSumTarget(root *TreeNode, k int, seen map[int]struct{}) bool {
    if root == nil {
        return false
    }

    if _, ok := seen[k - root.Val]; ok {
        return true
    }

    seen[root.Val] = struct{}{}

    return findSumTarget(root.Right, k, seen) || findSumTarget(root.Left, k, seen)
}
