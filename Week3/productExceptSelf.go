package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var product = make([]int, len(nums))
	var temp, i int = 1, 0
	for i = 0; i < len(nums); i++ {
		product[i] = temp
		temp = temp * nums[i]
	}
	temp = 1

	for i = i - 1; i >= 0; i-- {
		product[i] = product[i] * temp
		temp = temp * nums[i]
	}
	return product
}

func main() {
	nums := []int{1, 4, 0}
	// [1, 4, 4]
	// [0,0,4]
	fmt.Println(productExceptSelf(nums))
}
