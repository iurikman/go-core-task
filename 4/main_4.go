package main

import (
	"fmt"
)

func main() {
	s1 := []string{"aa", "bb", "cc"}
	s2 := []string{"aaa", "bbb", "ccc", "aa", "cc"}
	fmt.Println(filterSlice(s1, s2))
}

func filterSlice(slice1, slice2 []string) []string {
	m := make(map[string]struct{})
	for _, v := range slice1 {
		m[v] = struct{}{}
	}

	res := make([]string, 0, len(m))

	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}
	res = res[:len(res)]
	return res
}
