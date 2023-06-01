package product_of_array_exceptself

func ProductExceptSelf(nums []int) []int {
	l := make([]int, len(nums)+1)
	r := make([]int, len(nums)+1)
	r[len(nums)] = 1
	l[0] = 1
	for i := len(nums) - 1; i >= 0; i-- {
		r[i] = r[i+1] * nums[i]
	}
	for i := 1; i < len(nums)+1; i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	ans := make([]int, len(nums))
	for i := range nums {
		ans[i] = l[i] * r[i+1]
	}
	return ans
}
