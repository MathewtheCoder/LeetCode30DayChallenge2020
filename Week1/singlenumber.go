package main

import "fmt"

func singleNumber(nums []int) int {
	candidates := make(map[int]bool)
	for _, num := range nums {
		if candidates[num] {
			delete(candidates, num)
		} else {
			candidates[num] = true
		}
	}
	for Number := range candidates {
		return Number
	}
	return 0
}

// If we take XOR of zero and some bit, it will return that bit
// a ⊕ 0 = a
// If we take XOR of two same bits, it will return 0
// a ⊕ a = 0
// a ⊕ b ⊕ a = (a ⊕ a) ⊕ b = b
func singleNumberSolution(nums []int) int {
	a := 0
	for _, num := range nums {
		a ^= num
	}
	return a
}

func main() {
	nums := []int{2, 1, 2, 3, 1}

	fmt.Println(singleNumberSolution(nums))
}
