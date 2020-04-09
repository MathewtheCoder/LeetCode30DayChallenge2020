package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
	previousNumbers := make(map[int]bool)
	num := n
	var sum int = 0
	for {
		fmt.Println(math.Pow(float64(num%10), 2))
		sum += int(math.Pow(float64(num%10), 2))
		num = num / 10
		if num == 0 {
			fmt.Println(sum)
			fmt.Println("===================")
			// Return True if sum is 1.
			if sum == 1 {
				return true
			}
			if previousNumbers[sum] {
				return false
			} else {
				previousNumbers[sum] = true
			}
			num = sum
			sum = 0
		}
	}
}
func isHappyFast(n int) bool {
	seen := make(map[int]bool)
	return dfs(n, seen)
}

func dfs(n int, seen map[int]bool) bool {
	if seen[n] == true {
		return false
	}

	if n == 1 {
		return true
	}

	seen[n] = true

	return dfs(sumOfSquares(n), seen)
}

func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
func main() {
	num := 20
	isHappy(num)
}
