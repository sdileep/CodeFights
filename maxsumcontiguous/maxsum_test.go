package maxsumcontiguous_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/maxsumcontiguous"
)

var testCases = []struct {
	in       []int
	expected int
}{
	{[]int{-2, 2, 5, -11, 6}, 7},
	{[]int{-3, -2, -1, -4}, -1},
	{[]int{-3, 2, 1, -4}, 3},
}

func TestArrayMaxConsecutiveSum2(t *testing.T) {
	for _, test := range testCases {
		got := maxsumcontiguous.ArrayMaxConsecutiveSum2(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, want %d", got, test.expected)
		}
	}
}
