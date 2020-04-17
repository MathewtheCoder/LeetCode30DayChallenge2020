package main

import "fmt"

// 0(n^2)
func findMaxLengthBrute(nums []int) int {
	var i, j, maxLen int
	var counter = [2]int{0, 0}
	var arrLen = len(nums)
	for i = 0; i < arrLen; i++ {
		counter[0], counter[1] = 0, 0
		counter[nums[i]]++
		if maxLen > arrLen-i {
			// fmt.Println("Break", i)
			break
		}
		for j = i + 1; j < arrLen; j++ {
			counter[nums[j]]++
			if (counter[0] == counter[1]) && (maxLen < ((j + 1) - i)) {
				maxLen = (j + 1) - i
				// fmt.Println("Max", maxLen)
			}
		}
	}
	return maxLen
}
func findMaxLength(nums []int) int {
	var i, count, maxLen int
	var counter = make(map[int]int)
	counter[0] = -1
	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = count - 1
		} else {
			count = count + 1
		}
		if value, ok := counter[count]; ok {
			maxLen = Max(maxLen, i-value)
		} else {
			counter[count] = i
		}
		fmt.Println(counter)
	}
	return maxLen
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	nums := []int{1, 1, 1, 0, 0, 0}
	fmt.Println(findMaxLength(nums))
}
