package main

import (
	"math/rand"
	"fmt"
)

func main() {
	originalSlice := make([]int, 6)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(10)
	}
	fmt.Println("init slice originalSlice:")
	fmt.Println(originalSlice)

	fmt.Println("sliceExample:")
	fmt.Println(sliceExample(originalSlice))

	fmt.Println("add element 888 to init slice:")
	fmt.Println(addElement(originalSlice, 888))

	fmt.Println("copy init slice originalSlice:")
	copied := copySlice(originalSlice)
	fmt.Println(copied)

	fmt.Println("amended init slice:")
	originalSlice[1] = 333
	fmt.Println(originalSlice)

	fmt.Println("copied slice originalSlice:")
	fmt.Println(copied)

	fmt.Println("remove element with index 1:")
	fmt.Println(removeElement(originalSlice, 1))
}

func sliceExample(s []int) []int {
	result := make([]int, 0, len(s))
	for i := range s {
		if i%2 == 0 {
			result = append(result, i)
		}
		result = result[:len(result)]
	}

	return result
}

func addElement(s []int, num int) []int {
	s = append(s, num)
	return s
}

func copySlice(s []int) []int {
	result := make([]int, len(s))
	copy(result, s)
	return result
}

func removeElement(s []int, index int) []int {
	s = append(s[:index], s[index+1:]...)
	return s
}
