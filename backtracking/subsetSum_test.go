package backtracking_test

import (
	"reflect"
	"testing"

	"github.com/sunilgopinath/CodeFights/backtracking"
)

var testSubsetCases = []struct {
	arr      []int
	num      int
	expected [][]int
}{
	{
		[]int{1, 2, 3, 4, 5},
		5,
		[][]int{{1, 4}, {2, 3}, {5}},
	},
	{
		[]int{1, 1, 2, 2, 2, 5, 5, 6, 8, 8},
		9,
		[][]int{{1, 1, 2, 5}, {1, 2, 6}, {1, 8}, {2, 2, 5}},
	},
}

func TestSumSubsets(t *testing.T) {
	for _, test := range testSubsetCases {
		got := backtracking.SumSubsets(test.arr, test.num)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Got %v, expected %v from %v, %d", got, test.expected, test.arr, test.num)
		}
	}
}
