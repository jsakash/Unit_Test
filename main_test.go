package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {

	if Calculate(2) != 4 {
		t.Errorf("Expected 2 + 2 = 4 ")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
	}
	for _, tests := range tests {
		if output := Calculate(tests.input); output != tests.expected {
			t.Error("Test Failed: {} inputed, {} expected, {} received: {}", tests.input, tests.expected, output)
		}
	}
}

func TestMultiply(t *testing.T) {

	result := Multiply(2)
	if result%10 != 0 {
		t.Error("Expected multiple of ten ")
	}
}
