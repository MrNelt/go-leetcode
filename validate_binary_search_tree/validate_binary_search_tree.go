package validate_binary_searchtree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return solve(root, -math.MaxInt, math.MaxInt)
}

func solve(node *TreeNode, mini, maxi int) bool {
	if node == nil {
		return true
	}
	if node.Val <= mini {
		return false
	}

	if node.Val >= maxi {
		return false
	}
	return solve(node.Right, max(mini, node.Val), maxi) && solve(node.Left, mini, min(maxi, node.Val))
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
