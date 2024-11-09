package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	groupIdx := make(map[string]int)

	for _, str := range strs {
		letters := strings.Split(str, "")
		sort.Strings(letters)
		cpy := strings.Join(letters, "")

		if v, ok := groupIdx[cpy]; ok {
			res[v] = append(res[v], str)
		} else {
			groupIdx[cpy] = len(res)
			res = append(res, []string{str})
		}
	}

	return res
}
