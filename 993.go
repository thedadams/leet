/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x, y int) bool {
	pos := map[int]*TreeNode{
		x: nil,
		y: nil,
	}

	findDepths(nil, root, pos, 0)
	if pos[x] == nil || pos[y] == nil {
		return false
	}

	return pos[x] != pos[y] && pos[x].Val == pos[y].Val
}

func findDepths(parent, root *TreeNode, pos map[int]*TreeNode, depth int) {
	if root == nil {
		return
	}

	for i := range pos {
		if root.Val == i {
			if parent == nil {
				return
			}

			// This is a hack, but works because we are doing DFS.
			parent.Val = depth
			pos[i] = parent
		}
	}

	findDepths(root, root.Right, pos, depth+1)
	findDepths(root, root.Left, pos, depth+1)
}
