package main

import "fmt"

// func checkValidString(s string) bool {
// 	var arr []string
// 	for i := 0; i < len(s); i++ {
// 		if string(s[i]) == "(" || string(s[i]) == "*" {
// 			arr = append(arr, string(s[i]))
// 		} else if string(s[i]) == ")" {
// 			// fmt.Println(")")
// 			if len(arr) > 0 && (string(arr[len(arr)-1]) == "(" || string(arr[len(arr)-1]) == "*") {
// 				arr = arr[:len(arr)-1]
// 				// fmt.Println(i, arr, len(arr))
// 			} else {
// 				return false
// 			}
// 		}
// 		// fmt.Println(arr)
// 	}
// 	return len(arr) == 0
// }

func checkValidString(s string) bool {
	if len(s) == 0 {
		return true
	}
	var lo, hi int
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			lo++
			hi++
		} else if string(s[i]) == ")" {
			if lo > 0 {
				lo--
			}
			hi--
		} else {
			if lo > 0 {
				lo--
			}
			hi++
		}
		if hi < 0 {
			return false
		}
	}
	return lo == 0
}

func main() {
	str := "(*)"
	fmt.Println(checkValidString(str))
}
