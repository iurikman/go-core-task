package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestCrossElements(t *testing.T) {
	s1 := []int{1, 5, 2, 3, 6}
	s2 := []int{0, 2, 5}

	expected := []int{2, 5}

	res := crossElements(s1, s2)
	fmt.Println(res, expected)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("result: %v, expected: %v", res, expected)
	}
}
