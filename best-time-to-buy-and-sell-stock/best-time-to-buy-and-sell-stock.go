package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	ans := 0
	curPrice := prices[0]
	for _, tempPrice := range prices {
		if tempPrice <= curPrice {
			curPrice = tempPrice
		}
		ans = max(ans, tempPrice-curPrice)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
