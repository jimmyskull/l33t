package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	table := make(map[string][]string)
	for _, val := range strs {
		sval := SortString(val)
		table[sval] = append(table[sval], val)
	}
	var result [][]string
	for _, list := range table {
		result = append(result, list)
	}
	return result
}

func main() {
	values := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(values)
	fmt.Println(result)
}
