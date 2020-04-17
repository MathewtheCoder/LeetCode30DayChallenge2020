package main

import "fmt"

func stringShift(s string, shift [][]int) string {
	var strLen int = len(s)
	for _, arr := range shift {
		var noOfShifts int = arr[1] % strLen
		if arr[0] == 0 && arr[1] != 0 {
			var shifted string = s[:noOfShifts]     // Take first noOfShifts chars
			var destination string = s[noOfShifts:] // Take chars after noOfShifts
			s = destination + shifted
			// fmt.Println(strLen, noOfShifts, shifted, destination, s)
		} else if arr[0] == 1 && arr[1] != 0 {
			var destination string = s[(strLen - noOfShifts):]
			var shifted string = s[:(strLen - noOfShifts)]
			s = destination + shifted
			// fmt.Println(strLen, noOfShifts, shifted, destination, s)
		}
	}
	return s
}

func main() {
	s := "mathew"
	var shiftArr = [][]int{{0, 3}, {1, 2}, {1, 16}, {0, 1}}
	fmt.Println(stringShift(s, shiftArr))
}
