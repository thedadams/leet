import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	result := &[]int{math.MaxInt}[0]
	inOrder(root, new(*TreeNode), result)
	return *result
}

func inOrder(root *TreeNode, prev **TreeNode, minSoFar *int) {
	if root != nil {
		inOrder(root.Left, prev, minSoFar)
		if *prev != nil {
			*minSoFar = min(*minSoFar, root.Val-(*prev).Val)
		}
		*prev = root
		inOrder(root.Right, prev, minSoFar)
	}
}

