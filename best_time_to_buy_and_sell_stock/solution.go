package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	buy := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else if prices[i]-buy > profit {
			profit = prices[i] - buy
		}
	}

	return profit
}
