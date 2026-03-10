/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    sums, counts := make([]float64, 0), make([]int, 0)
    sums, counts = averageOfLevelsAtLevel(root, 0, sums, counts)
    for i := range sums {
        sums[i] /= float64(counts[i])
    }

    return sums
}

func averageOfLevelsAtLevel(root *TreeNode, level int, sums []float64, counts []int) ([]float64, []int) {
    if root == nil {
        return sums, counts
    }

    sums, counts = averageOfLevelsAtLevel(root.Left, level + 1, sums, counts)
    sums, counts = averageOfLevelsAtLevel(root.Right, level + 1, sums, counts)
    
    for len(sums) <= level {
        sums = append(sums, 0)
        counts = append(counts, 0)
    }

    sums[level] += float64(root.Val)
    counts[level]++

    return sums, counts
}