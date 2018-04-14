package arcade_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/arcade"
)

var testAdjacentCases = []struct {
	in       []int
	expected int
}{
	{[]int{3, 6, -2, -5, 7, 3}, 21},
	{[]int{-1, -2}, 2},
	{[]int{5, 1, 2, 3, 1, 4}, 6},
	{[]int{-23, 4, -3, 8, -12}, -12},
}

func TestAdjacentElements(t *testing.T) {
	for _, test := range testAdjacentCases {
		got := arcade.AdjacentElementsProduct(test.in)
		if got != test.expected {
			t.Fatalf("Got %d, want %d", got, test.expected)
		}
	}
}
