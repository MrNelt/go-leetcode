package lowest_common_ancestor_of_a_binary_search_tree

type Stack struct {
	stack []*TreeNode
}

func NewStack() Stack {
	return Stack{[]*TreeNode{}}
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) MustTop() *TreeNode {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) MustPop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack) MustPush(i *TreeNode) {
	s.stack = append(s.stack, i)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	stackP := NewStack()
	stackQ := NewStack()
	dfs(root, p.Val, &stackP)
	dfs(root, q.Val, &stackQ)
	for stackP.MustTop() != stackQ.MustTop() {
		if stackP.Len() > stackQ.Len() {
			stackP.MustPop()
		} else {
			stackQ.MustPop()
		}
	}
	return stackP.MustTop()
}

func dfs(node *TreeNode, val int, s *Stack) {
	s.MustPush(node)
	if node.Val == val {
		return
	} else if val > node.Val {
		dfs(node.Right, val, s)
	} else {
		dfs(node.Left, val, s)
	}
}
