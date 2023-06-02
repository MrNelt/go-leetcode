package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	dfs(root)
	return root
}

func dfs(node *TreeNode) {
	if node != nil {
		temp := node.Left
		node.Left = node.Right
		node.Right = temp
		dfs(node.Left)
		dfs(node.Right)
	}
}
