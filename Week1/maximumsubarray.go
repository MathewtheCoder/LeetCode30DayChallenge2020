package main

import (
	"fmt"
	"math"
)

// My Solution
func maxSubArray(nums []int) int {
	max_so_far := -math.MaxUint32 - 1
	max_ending_here := 0
	for _, num := range nums {
		max_ending_here = max_ending_here + num
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}
		if max_ending_here < 0 {
			max_ending_here = 0
		}
		fmt.Println("max ending here", max_ending_here)
		fmt.Println("max so far", max_so_far)
	}
	return max_so_far
}

// Best solution for this
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxSubArrayFast(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sum := 0
	msum := nums[0]

	for _, n := range nums {
		sum = max(sum+n, n)
		msum = max(msum, sum)
	}

	return msum

}

// 1 -3 5 -10 1
func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{5, 1, -3, 4, 7, -5, 2}
	fmt.Println(maxSubArray(nums))
}
