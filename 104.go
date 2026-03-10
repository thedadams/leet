/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    var depth int
    if root != nil {
        depth = max(depth, 1+maxDepth(root.Right))
        depth = max(depth, 1+maxDepth(root.Left))
    }

    return depth
}