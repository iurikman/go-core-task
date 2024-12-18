package main

import "fmt"

func main() {
	s := StringIntMap{make(map[string]int)}
	s.Add("one", 1)
	s.Add("two", 2)
	s.Add("three", 3)
	fmt.Println(s.m)

	s.Remove("three")
	fmt.Println(s.m)

	newMap := s.Copy()
	fmt.Println(newMap)
	fmt.Println(s.m)

	fmt.Println(s.Exists("one"))
	fmt.Println(s.Exists("random"))
}

type StringIntMap struct {
	m map[string]int
}

func (s *StringIntMap) Add(key string, value int) {
	s.m[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.m, key)
}

func (s *StringIntMap) Copy() map[string]int {
	res := make(map[string]int)
	for key, value := range s.m {
		res[key] = value
	}
	return res
}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.m[key]
	if !ok {
		return false
	}

	return true
}

func (s *StringIntMap) Get(key string) (int, bool) {
	value, ok := s.m[key]
	if !ok {
		return 0, false
	}

	return value, true
}
