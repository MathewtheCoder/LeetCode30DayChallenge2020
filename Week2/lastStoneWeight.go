package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	// fmt.Println(stones)
	l := len(stones)
	for len(stones) > 1 {
		first, second := stones[l-1], stones[l-2]
		if first == second {
			stones = stones[:l-2]
			l = l - 2
		} else {
			stones = stones[:l-2]
			stones = append(stones, first-second)
			l = l - 1
		}
		sort.Ints(stones)
		fmt.Println(stones)

	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

func main() {
	// nums := []int{2, 7, 4, 1, 8, 1}
	nums := []int{10, 4, 2, 10}
	fmt.Println(lastStoneWeight(nums))
}
