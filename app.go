package main

import (
	"fmt"
	largest_rectangle_inhistogram "go-leetcode/largest_rectangle_in_histogram"
	"go-leetcode/reorder_list"
)

func main() {
	fmt.Println(largest_rectangle_inhistogram.LargestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	h := reorder_list.NewList(1, 2, 3, 4)
	h.Print()
	h.Print()
	h.Print()
	reorder_list.ReorderList(h)
}
