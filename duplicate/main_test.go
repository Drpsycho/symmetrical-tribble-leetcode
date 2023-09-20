package duplicate

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	// Test case 1: nums contains duplicates; expected output: true
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 2}
	result1 := containsDuplicate(nums1)
	if !result1 {
		t.Errorf("Test case 1 failed: Expected true, got false")
	}

	// Test case 2: nums does not contain duplicates; expected output: false
	nums2 := []int{1, 2, 3, 4, 1, 6, 7, 8, 9}
	result2 := containsDuplicate(nums2)
	if !result2 {
		t.Errorf("Test case 2 failed: Expected false, got true")
	}

	// Test case 3: nums is empty; expected output: false
	nums3 := []int{}
	result3 := containsDuplicate(nums3)
	if result3 {
		t.Errorf("Test case 3 failed: Expected false, got true")
	}

	// Test case 4: nums has only one element; expected output: false
	nums4 := []int{1}
	result4 := containsDuplicate(nums4)
	if result4 {
		t.Errorf("Test case 4 failed: Expected false, got true")
	}
}
