package main

import (
	"testing"
	"fmt"
)

func TestGetType(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected string
	}{
		{12, "int"},
		{0o34, "int"},
		{0x56, "int"},
		{7.8, "float64"},
		{"string", "string"},
		{true, "bool"},
		{complex64(9), "complex64"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			getType(test.input)
			if test.expected != fmt.Sprintf("%T", test.input) {
				t.Errorf("expected: %T, got: %v", test.expected, test.input)
			}
		})

	}
}

func TestCreateString(t *testing.T) {
	decimal := 12
	octal := 0o34
	hexVal := 0x56
	float := 7.8
	str := "test"
	boolean := true
	compl := complex64(2)

	expected := "1234567.800000testtrue(2+0i)"
	result := createString(decimal, octal, hexVal, float, str, boolean, compl)

	fmt.Print(result)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestConvertToRune(t *testing.T) {
	input := "test1"
	expected := []rune{'t', 'e', 's', 't', '1'}

	result := convertToRune(input)

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i, r := range expected {
		if result[i] != r {
			t.Errorf("At index %d: expected %c, got %c", i, r, result[i])
		}
	}
}
