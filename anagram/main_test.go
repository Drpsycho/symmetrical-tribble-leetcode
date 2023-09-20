package anagram

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	// Test case 1: nums contains duplicates; expected output: true{
	first := "anagram"
	second := "nagaram"
	result1 := isAnagram(first, second)
	if !result1 {
		t.Errorf("Test case 1 failed: Expected true, got false")
	}

	first = "aa"
	second = "bb"
	result1 = isAnagram(first, second)
	if result1 {
		t.Errorf("Test case 1 failed: Expected true, got false")
	}
}
