package q2588

import (
	"testing"
)

func TestBeautifulSubarrays(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int64
	}{
		{[]int{4, 3, 1, 2, 4}, 2},
		{[]int{1, 10, 4}, 0},
	}

	for _, tc := range testCases {
		result := beautifulSubarrays(tc.nums)
		if result != tc.expected {
			t.Errorf("For input %v, expected %d but got %d", tc.nums, tc.expected, result)
		}
	}
}
