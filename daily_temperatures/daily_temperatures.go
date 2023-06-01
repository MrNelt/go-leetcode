package daily_temperatures

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

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	s := NewStack()
	for i := range temperatures {
		for s.Len() > 0 && temperatures[i] > temperatures[s.MustTop()] {
			ans[s.MustTop()] = i - s.MustTop()
			s.MustPop()
		}
		s.MustPush(i)
	}
	return ans
}
