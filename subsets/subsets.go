package subsets

type Stack struct {
	stack []int
}

func NewStack() Stack {
	return Stack{[]int{}}
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) MustTop() int {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) MustPop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack) MustPush(i int) {
	s.stack = append(s.stack, i)
}

func subsets(nums []int) [][]int {
	ans := [][]int{}
	stack := NewStack()
	solve(0, nums, &stack, &ans)
	return ans
}

func solve(i int, nums []int, stack *Stack, ans *[][]int) {
	if i == len(nums) {
		temp := []int{}
		for _, v := range stack.stack {
			temp = append(temp, v)
		}
		*ans = append(*ans, temp)
		return
	}
	stack.MustPush(nums[i])
	solve(i+1, nums, stack, ans)
	stack.MustPop()
	solve(i+1, nums, stack, ans)
}
