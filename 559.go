/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }

    var md int
    for _, c := range root.Children {
        md = max(md, maxDepth(c))
    }

    return md + 1
}