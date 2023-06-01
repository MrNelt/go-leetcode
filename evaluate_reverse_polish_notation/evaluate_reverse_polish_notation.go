package evaluate_reverse_polish_notation

import (
	"strconv"
	"unicode"
)

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

func EvalRPN(tokens []string) int {
	s := NewStack()
	for _, v := range tokens {
		if unicode.IsDigit(rune(v[len(v)-1])) {
			num, _ := strconv.Atoi(v)
			s.MustPush(num)
		} else {
			num1 := s.MustTop()
			s.MustPop()
			num2 := s.MustTop()
			s.MustPop()
			if v == "+" {
				s.MustPush(num1 + num2)
			} else if v == "-" {
				s.MustPush(num2 - num1)
			} else if v == "*" {
				s.MustPush(num1 * num2)
			} else if v == "/" {
				s.MustPush(num2 / num1)
			}
		}
	}
	return s.MustTop()
}
