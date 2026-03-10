/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    if root == nil {
        return nil
    }

    vals := []int{root.Val}
    for _, child := range root.Children {
        vals = append(vals, preorder(child)...)
    }

    return vals
}