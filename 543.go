/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    _, diam := depth(root)
    return diam
}

func depth(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }

    leftDepth, leftDiam := depth(root.Left)
    rightDepth, rightDiam := depth(root.Right)

    return max(leftDepth, rightDepth) + 1, max(leftDiam, max(rightDiam, leftDepth + rightDepth))
}