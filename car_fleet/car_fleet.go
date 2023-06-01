package car_fleet

import (
	"fmt"
	"sort"
)

type Stack struct {
	stack []float64
}

func NewStack() Stack {
	return Stack{[]float64{}}
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) MustTop() float64 {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) MustPop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack) MustPush(i float64) {
	s.stack = append(s.stack, i)
}

func CarFleet(target int, position []int, speed []int) int {

	time := make([]float64, len(position))

	for i := range position {
		time[i] = float64((target - position[i])) / float64(speed[i])
	}

	sort.Slice(time, func(i, j int) bool {
		return time[i] < time[j]
	})

	fmt.Println(time)

	s := NewStack()

	for _, v := range time {
		for s.Len() > 0 && v >= s.MustTop() {
			s.MustPop()
		}
		s.MustPush(v)
	}
	return s.Len()
}
