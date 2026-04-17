import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1, leafs2 := make([]int, 0), make([]int, 0)
	leafSequence(root1, &leafs1)
	leafSequence(root2, &leafs2)
	return slices.Equal(leafs1, leafs2)
}

func leafSequence(root *TreeNode, leafs *[]int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*leafs = append(*leafs, root.Val)
		} else {
			leafSequence(root.Left, leafs)
			leafSequence(root.Right, leafs)
		}
	}
}
