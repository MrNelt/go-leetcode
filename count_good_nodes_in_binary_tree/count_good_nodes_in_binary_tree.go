package count_good_nodes_in_binarytree

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return solve(root, root.Val)
}

func solve(node *TreeNode, maxElemFromRoot int) int {
	if node == nil {
		return 0
	}
	maxElemFromRoot = max(maxElemFromRoot, node.Val)
	count := 0
	if maxElemFromRoot == node.Val {
		count = 1
	}
	return count + solve(node.Left, maxElemFromRoot) + solve(node.Right, maxElemFromRoot)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
