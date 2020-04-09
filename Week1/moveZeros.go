package main

import "fmt"

func moveZeroesNotInline(nums []int) []int {
	var result = make([]int, len(nums))
	counter := 0
	fmt.Println(result)

	for _, num := range nums {
		if num != 0 {
			result[counter] = num
			counter++
		}
	}
	return result
}
func moveZeros(nums []int) {
	l := len(nums)
	i := 0
	b := nums[:0]
	for i < l {
		if nums[i] == 0 {
			nums = append(b, append(nums[i+1:], 0)...)
			l--
		} else {
			i++
		}
	}
}

func moveZeroesFast(nums []int) {
	var i, zeros, ln, tmp int

	ln = len(nums)
	for i < ln {
		if nums[i] == 0 {
			zeros++
			i++
			break
		}
		i++
	}
	for i < ln {
		if nums[i] == 0 {
			zeros++
		} else {
			tmp = nums[i]
			nums[i] = nums[i-zeros]
			nums[i-zeros] = tmp
		}
		i++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	// fmt.Println(moveZeroesNotInline(nums))
	moveZeros(nums)
	// moveZeroesFast(nums)
	fmt.Println(nums)
}
