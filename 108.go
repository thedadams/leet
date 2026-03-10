/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    centerIsh := len(nums) / 2
    root := &TreeNode{Val: nums[centerIsh]}
    root.Left = sortedArrayToBST(nums[:centerIsh])
    root.Right = sortedArrayToBST(nums[centerIsh+1:])

    return root
}