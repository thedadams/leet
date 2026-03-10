/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    lh, rh := depth(root.Left, 0), depth(root.Right, 0)
    if lh == rh {
        return (1 << rh) + countNodes(root.Right)
    }
    return (1 << rh) + countNodes(root.Left)
}

func depth(root *TreeNode, count int) int {
    if root == nil {
        return count
    }

    return depth(root.Left, count + 1)
}
