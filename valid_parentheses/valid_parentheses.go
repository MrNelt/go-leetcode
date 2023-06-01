package valid_parentheses

type Stack struct {
	stack []rune
}

func NewStack() Stack {
	return Stack{[]rune{}}
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) MustTop() rune {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) MustPop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack) MustPush(i rune) {
	s.stack = append(s.stack, i)
}

func isValid(s string) bool {
	stack := NewStack()

	for _, v := range s {
		if v == '(' {
			stack.MustPush(v)
		} else if v == '{' {
			stack.MustPush(v)
		} else if v == '[' {
			stack.MustPush(v)
		} else {
			if v == ')' && (stack.Len() == 0 || stack.MustTop() != '(') {
				return false
			} else if v == '}' && (stack.Len() == 0 || stack.MustTop() != '{') {
				return false
			} else if v == ']' && (stack.Len() == 0 || stack.MustTop() != '[') {
				return false
			}
			stack.MustPop()
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}
