package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	return process(S) == process(T)
}
func process(S string) string {
	var ret string = ""
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "#" && len(ret) == 1 {
			ret = ""
		} else if string(S[i]) == "#" && ret != "" {
			ret = ret[:len(ret)-1]
		} else if string(S[i]) != "#" {
			ret = ret + string(S[i])
		}
	}
	return ret
}
func main() {
	// S := "ab#c"
	// T := "ad#c"
	S := "bxj##tw"
	T := "bxo#j##tw"
	fmt.Println(backspaceCompare(S, T))
}
