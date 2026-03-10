/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    tilt, _, _ := tiltAndSums(root)
    return tilt
}

func tiltAndSums(root *TreeNode) (int, int, int) {
    if root == nil {
        return 0, 0, 0
    }

    leftTilt, leftSum, rightSum := tiltAndSums(root.Left)
    leftSum += rightSum
    if root.Left != nil {
        leftSum += root.Left.Val
    }

    rightTilt, leftRightSum, rightSum := tiltAndSums(root.Right)
    rightSum += leftRightSum
    if root.Right != nil {
        rightSum += root.Right.Val
    }

    return rightTilt+ leftTilt + int(math.Abs(float64(leftSum - rightSum))), leftSum, rightSum
}