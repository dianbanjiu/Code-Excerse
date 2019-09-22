package leetcode

func MaxProfit2(prices []int) int {
	peek := 0
	valley := 0
	maxprofit := 0
	i := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peek = prices[i]
		maxprofit += peek - valley
	}
	return maxprofit
}