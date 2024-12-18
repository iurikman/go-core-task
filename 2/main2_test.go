package main

import (
	"testing"
	"reflect"
)

func TestSliceExample(t *testing.T) {
	initSlice := []int{0, 1, 2, 3, 4, 5, 6}
	expected := []int{0, 2, 4, 6}

	result := sliceExample(initSlice)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Slice example failed, expected %v, got %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	initSlice := []int{0, 1, 2, 3, 4, 5, 6}
	expected := []int{0, 1, 2, 3, 4, 5, 6}

	result := copySlice(initSlice)
	initSlice[0] = 999
	if !reflect.DeepEqual(result, expected) || reflect.DeepEqual(initSlice, result) {
		t.Errorf("Slice example failed, expected %v, got %v", expected, result)
	}
}

func TestRemoveElement(t *testing.T) {
	initSlice := []int{0, 1, 2, 3, 4, 5, 6}
	expected := []int{0, 2, 3, 4, 5, 6}

	result := removeElement(initSlice, 1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Slice example failed, expected %v, got %v", expected, result)
	}
}
