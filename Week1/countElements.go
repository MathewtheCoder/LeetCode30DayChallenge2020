package main

import (
	"fmt"
)

func countElements(arr []int) int {
	uniqueNumbers := make(map[int]float64)
	var count float64
	// place all the elements in the hashmap
	for _, num := range arr {
		uniqueNumbers[num]++
	}
	// Loop through and check if num + 1 is present in hash map
	for key, itemCount := range uniqueNumbers {
		if uniqueNumbers[key+1] > 0 {
			count = count + itemCount
		}
	}
	return int(count)
}
func main() {
	nums := []int{1, 1, 2}
	fmt.Println(countElements(nums))
}
