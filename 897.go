/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	return increasingBSTWithTail(root, nil)
}

func increasingBSTWithTail(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}

	newRoot := increasingBSTWithTail(root.Left, root)
	root.Left = nil
	root.Right = increasingBSTWithTail(root.Right, tail)
	return newRoot
}
