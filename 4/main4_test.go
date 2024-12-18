package main

import (
	"testing"
	"reflect"
)

func TestFilterSlice(t *testing.T) {
	slice1 := []string{"aa", "bb", "cc"}
	slice2 := []string{"aa", "bb", "cccc", "dddd"}

	expected := []string{"aa", "bb"}

	res := filterSlice(slice1, slice2)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Got %v, expected %v", res, expected)
	}
}
