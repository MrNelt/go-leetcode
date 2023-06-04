package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	cache := make(map[[40]int]bool)
	set := [40]int{}
	solve(0, target, candidates, &set, cache, &ans)
	return ans
}

func solve(tempSum, target int, canditates []int, set *[40]int, cache map[[40]int]bool, ans *[][]int) {
	if tempSum == target && !cache[*set] {
		temp := []int{}
		for i := range canditates {
			for j := 0; j < set[i]; j++ {
				temp = append(temp, canditates[i])
			}
		}
		*ans = append(*ans, temp)
		cache[*set] = true
		return
	}
	if tempSum >= target {
		return
	}
	for i, v := range canditates {
		set[i]++
		tempSum += v
		solve(tempSum, target, canditates, set, cache, ans)
		tempSum -= v
		set[i]--
	}
}
