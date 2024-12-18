package main

import (
	"math/rand"
	"fmt"
)

func main() {
	s1 := createSlice(6)
	s2 := createSlice(3)
	fmt.Println(s1, "\n", s2)
	fmt.Println(crossElements(s1, s2))
}

func crossElements(s1, s2 []int) interface{} {
	m := make(map[int]struct{})

	for _, v := range s1 {
		m[v] = struct{}{}
	}

	res := make([]int, 0)

	for _, v := range s2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}

	if len(res) != 0 {
		return res
	}

	return false

}

func createSlice(len int) []int {
	s := make([]int, len)
	for i := range s {
		s[i] = rand.Intn(10)
	}
	return s
}
