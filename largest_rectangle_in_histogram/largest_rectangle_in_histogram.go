package largest_rectangle_inhistogram

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

func LargestRectangleArea(heights []int) int {
	minRightStack := NewStack()
	minLeftStack := NewStack()

	minLeft := make([]int, len(heights))
	minRight := make([]int, len(heights))

	for i := range heights {
		for minRightStack.Len() > 0 && heights[i] < heights[minRightStack.MustTop()] {
			minRight[minRightStack.MustTop()] = i - 1
			minRightStack.MustPop()
		}
		minRightStack.MustPush(i)
	}

	for minRightStack.Len() > 0 {
		minRight[minRightStack.MustTop()] = len(heights) - 1
		minRightStack.MustPop()
	}

	for i := len(heights) - 1; i >= 0; i-- {
		for minLeftStack.Len() > 0 && heights[i] < heights[minLeftStack.MustTop()] {
			minLeft[minLeftStack.MustTop()] = i + 1
			minLeftStack.MustPop()
		}
		minLeftStack.MustPush(i)
	}

	for minLeftStack.Len() > 0 {
		minLeft[minLeftStack.MustTop()] = 0
		minLeftStack.MustPop()
	}

	ans := 0
	for i := range heights {
		// fmt.Println(heights[i], heights[i]*(minRight[i]-minLeft[i]+1))
		ans = max(ans, heights[i]*(minRight[i]-minLeft[i]+1))
	}
	// fmt.Println(heights)
	// fmt.Println(minRight)
	// fmt.Println(minLeft)

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
