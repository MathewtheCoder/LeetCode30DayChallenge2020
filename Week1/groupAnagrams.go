package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramGroupsMap := make(map[string][]string)
	var anagramGroups [][]string
	for _, str := range strs {
		var sortedString string = sortString(str)
		anagramGroupsMap[sortedString] = append(anagramGroupsMap[sortedString], str)
	}
	for _, arr := range anagramGroupsMap {
		anagramGroups = append(anagramGroups, arr)
	}
	return anagramGroups
}
func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func main() {
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	errorArr := []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	fmt.Println(groupAnagrams(errorArr))
}
