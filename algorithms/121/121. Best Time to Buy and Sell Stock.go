package algorithms

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var profit int
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			tmp := prices[i] - min
			if tmp > profit {
				profit = tmp
			}
		}
	}
	return profit
}
