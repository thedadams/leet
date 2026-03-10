/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    _, result := search(root, -1, math.MaxInt)
    return result
}

func search(root *TreeNode, prev, result int) (int, int) {
    if root == nil {
        return prev, result
    }

    prev, result = search(root.Left, prev, result)
    if prev != -1 {
        result = min(result, root.Val - prev)
    }

    return search(root.Right, root.Val, result)
}
