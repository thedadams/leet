/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    m := make(map[int]int)
    modeCount := findModeWithCounter(root, m)
    
    var result []int
    for i, c := range m {
        if c == modeCount {
            result = append(result, i)
        }
    }

    return result
}

func findModeWithCounter(root *TreeNode, counters map[int]int) int {
    if root == nil {
        return 0
    }

    leftModeCount := findModeWithCounter(root.Left, counters)
    rightModeCount := findModeWithCounter(root.Right, counters)
    
    counters[root.Val]++
    thisNodeCount := counters[root.Val]

    return max(leftModeCount, max(rightModeCount, thisNodeCount))
}