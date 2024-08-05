package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	text := "Hello world! Hello"
	expected := map[string]int{"hello": 2, "world": 1}
	result := countWords(text)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCheckPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal, Panama!", true},
		{"racecar", true},
		{"hello", false},
		{"Was it a car or a cat I saw?", true},
		{"No lemon, no melon", true},
		{"", true},
	}

	for _, test := range tests {
		result := checkPalindrome(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
