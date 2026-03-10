/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    if root == nil {
        return nil
    }

    vals := make([]int, 0, len(root.Children)+1)
    for _, c := range root.Children {
        vals = append(vals, postorder(c)...)
    }

    return append(vals, root.Val)
}