/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
   if root == nil {
    return nil
   } 

   result := inorderTraversal(root.Left)
   result = append(result, root.Val)
   result = append(result, inorderTraversal(root.Right)...)

   return result
}