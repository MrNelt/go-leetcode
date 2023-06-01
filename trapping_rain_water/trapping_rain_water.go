package trapping_rain_water

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

func Trap(height []int) int {
	ans := 0
	ge := make([]int, len(height))
	s := NewStack()
	for i := range height {
		for s.Len() > 0 && height[s.MustTop()] <= height[i] {
			if height[s.MustTop()] != 0 {
				ge[s.MustTop()] = i
			} else {
				ge[s.MustTop()] = -1
			}
			s.MustPop()
		}
		s.MustPush(i)
	}
	ge[s.MustTop()] = -1
	value := s.MustTop()
	s.MustPop()
	for s.Len() > 0 {
		ge[s.MustTop()] = value
		value = s.MustTop()
		s.MustPop()
	}
	l := 0
	for l < len(height) && ge[l] == -1 {
		l++
	}
	if l >= len(height) {
		return 0
	}
	r := ge[l]
	for l < len(height) && ge[l] != -1 {
		t := min(height[l], height[r])
		for l != r {
			ans += max(t-height[l], 0)
			l++
		}
		r = ge[l]
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
