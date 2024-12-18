package main

import (
	"testing"
	"reflect"
)

func TestAdd(t *testing.T) {
	initMap := StringIntMap{make(map[string]int)}
	initMap.m["one"] = 1
	initMap.m["two"] = 2
	initMap.m["three"] = 3

	expect := 4
	initMap.Add("four", 4)

	if initMap.m["four"] != expect {
		t.Errorf("Expected %d, got %d", expect, initMap.m["four"])
	}
}

func TestRemove(t *testing.T) {
	initMap := StringIntMap{make(map[string]int)}
	initMap.m["one"] = 1
	initMap.m["two"] = 2
	initMap.m["three"] = 3

	initMap.Remove("three")
	expect := 0

	if initMap.m["three"] != expect {
		t.Errorf("expected %d, got %d", expect, initMap.m["three"])
	}
}

func TestCopy(t *testing.T) {
	initMap := StringIntMap{make(map[string]int)}
	initMap.m["one"] = 1
	initMap.m["two"] = 2
	initMap.m["three"] = 3

	expect := map[string]int{"one": 1, "two": 2, "three": 3}

	result := initMap.Copy()

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestExist(t *testing.T) {
	initMap := StringIntMap{make(map[string]int)}
	initMap.m["one"] = 1
	initMap.m["two"] = 2
	initMap.m["three"] = 3

	expect := true
	res := initMap.Exists("three")

	if res != expect {
		t.Errorf("Expected %v, got %v", expect, res)
	}
}

func TestGet(t *testing.T) {
	initMap := StringIntMap{make(map[string]int)}
	initMap.m["one"] = 1
	initMap.m["two"] = 2
	initMap.m["three"] = 3

	expect := 3

	res, ex := initMap.Get("three")
	if res != expect {
		t.Errorf("Expected %v, got %v", expect, res)
	}

	if ex != true {
		t.Errorf("Expected %v, got %v", true, res)
	}

	expect = 0
	res, ex = initMap.Get("four")
	if res != expect {
		t.Errorf("Expected %v, got %v", expect, res)
	}

	if ex != false {
		t.Errorf("Expected %v, got %v", true, res)
	}
}
