package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	var i, profit, buy, sell int
	for i < n-1 {
		for (i < n-1) && (prices[i+1] <= prices[i]) {
			i++
		}
		if i == n-1 {
			break
		}
		buy = i
		i++
		for (i < n) && (prices[i] >= prices[i-1]) {
			i++
		}
		i--
		sell = i
		profit = profit + (prices[sell] - prices[buy])
	}
	return profit
}
func maxProfitFast(prices []int) int {
	max := 0
	for i, p := range prices {
		if i == len(prices)-1 {
			break
		}
		fmt.Println("Here")
		if prices[i+1] > p {
			max += prices[i+1] - p
		}
	}
	return max
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	// price2 := []int{1, 2, 3, 4, 5}
	// prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfitFast(prices))
	// fmt.Println(maxProfit(price2))
	// fmt.Println(maxProfit(prices3))
}
