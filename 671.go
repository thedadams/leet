/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil {
		return -1
	}

	m := math.MaxInt
    if root.Left.Val != root.Right.Val {
        m = max(root.Left.Val, root.Right.Val)
    }
	if r := findSecondMinimumVal(root.Right, root.Val, m); r > root.Val {
		m = min(r, m)
	}
	if l := findSecondMinimumVal(root.Left, root.Val, m); l > root.Val {
		m = min(l, m)
	}

	if m == math.MaxInt {
		return -1
	}
	return m
}

func findSecondMinimumVal(root *TreeNode, absMin, minSoFar int) int {
	if root == nil || root.Left == nil {
		return -1
	}
	if root.Val >= minSoFar {
		return minSoFar
	}

	if m := max(root.Left.Val, root.Right.Val); m > root.Val {
		minSoFar = min(minSoFar, m)
	}

	if l := findSecondMinimumVal(root.Left, absMin, minSoFar); l != -1 {
		minSoFar = min(minSoFar, l)
	}
	if r := findSecondMinimumVal(root.Right, absMin, minSoFar); r != -1 {
		minSoFar = min(minSoFar, r)
	}
	return minSoFar
}