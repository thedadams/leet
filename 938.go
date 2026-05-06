/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}

	var (
		result  int
		inRange = true
	)
	if root.Val >= low {
		result += rangeSumBST(root.Left, low, high)
	} else {
		inRange = false
	}
	if root.Val <= high {
		result += rangeSumBST(root.Right, low, high)
	} else {
		inRange = false
	}

	if inRange {
		result += root.Val
	}

	return result
}

