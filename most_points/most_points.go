package most_points

func mostPoints(questions [][]int) int64 {

	n := len(questions)
	dp := make([]int, n)
	dp[n-1] = questions[n-1][0]
	for i := n - 2; i >= 0; i-- {
		jump := i + questions[i][1] + 1
		temp := questions[i][0]
		if jump < n {
			temp += dp[jump]
		}
		dp[i] = max(temp, dp[i+1])
	}

	return int64(dp[0])
}

func max(i, temp int) int {
	if i > temp {
		return i
	}
	return temp
}
