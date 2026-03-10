/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
   if root == nil {
    return 0
   }

   if root.Left == nil && root.Right == nil {
    return 0
   }

    var result int
   if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
    result = root.Left.Val
   }

   return result + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}