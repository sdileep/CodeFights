package basic_test

import (
	"testing"

	"github.com/sunilgopinath/codefights/basic"
)

var testSumCases = []struct {
	arr      []int
	qs       [][]int
	expected int
}{
	{
		[]int{3, 0, -2, 6, -3, 2},
		[][]int{{0, 2}, {2, 5}, {0, 5}},
		10,
	},
	{
		[]int{-1000},
		[][]int{{0, 0}},
		999999007,
	},
	{
		[]int{34, 19, 21, 5, 1, 10, 26, 46, 33, 10},

		[][]int{{3, 7},
			{3, 4},
			{3, 7},
			{4, 5},
			{0, 5}},
		283,
	},
}

func TestSumInRange(t *testing.T) {
	for _, test := range testSumCases {
		got := basic.SumInRange(test.arr, test.qs)
		if got != test.expected {
			t.Fatalf("Got %d, want %d", got, test.expected)
		}
	}
}
