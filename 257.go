/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        return []string{fmt.Sprintf("%d", root.Val)}
    }

    paths := binaryTreePaths(root.Left)
    paths = append(paths, binaryTreePaths(root.Right)...)

    for i, p := range paths {
        paths[i] = fmt.Sprintf("%d->%s", root.Val, p)
    }

    return paths
}